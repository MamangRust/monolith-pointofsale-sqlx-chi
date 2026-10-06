package role_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	apicache "github.com/MamangRust/monolith-pointofsale-apigateway/cache/gateway_cache"
	rolehandler "github.com/MamangRust/monolith-pointofsale-apigateway/handler"
	pbrole "github.com/MamangRust/monolith-pointofsale-pb/role"
	"github.com/MamangRust/monolith-pointofsale-pkg/logger"
	role_cache "github.com/MamangRust/monolith-pointofsale-role/cache"
	role_handler "github.com/MamangRust/monolith-pointofsale-role/handler"
	"github.com/MamangRust/monolith-pointofsale-role/repository"
	"github.com/MamangRust/monolith-pointofsale-role/service"
	"github.com/MamangRust/monolith-pointofsale-shared/cache"
	"github.com/MamangRust/monolith-pointofsale-shared/domain/requests"
	response_api "github.com/MamangRust/monolith-pointofsale-shared/mapper/response/api"
	"github.com/MamangRust/monolith-pointofsale-shared/observability"
	pb "github.com/MamangRust/monolith-pointofsale-apigateway/pbcompat"
	tests "github.com/MamangRust/monolith-pointofsale-test"
	"github.com/go-chi/chi/v5"

	"github.com/MamangRust/monolith-pointofsale-apigateway/apierror"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RoleApiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
	echo        chi.Router
	grpcServer  *grpc.Server
	conn        *grpc.ClientConn
	roleID      int
}

func (s *RoleApiTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := pgxpool.New(s.ts.Ctx, s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)

	queries, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	repos := repository.NewRepositories(queries)

	log, _ := logger.NewLogger("test", nil)
	obs, _ := observability.NewObservability("test", log)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)
	mencache := role_cache.NewMencache(cacheStore)

	roleService := service.NewService(&service.Deps{
		Repositories:  repos,
		Logger:        log,
		Mencache:      mencache,
		Observability: obs,
	})

	// Start internal gRPC Server for Role module
	roleHandlerGrpc := role_handler.NewHandler(&role_handler.Deps{
		Service: roleService,
		Logger:  log,
	})
	server := grpc.NewServer()
	pbrole.RegisterRoleQueryServiceServer(server, roleHandlerGrpc.Role)
	pbrole.RegisterRoleCommandServiceServer(server, roleHandlerGrpc.RoleCommand)
	s.grpcServer = server

	lis, err := net.Listen("tcp", "localhost:0")
	s.Require().NoError(err)

	go func() {
		_ = server.Serve(lis)
	}()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn

	// Setup Echo and API Handler
	e := chi.NewRouter()
	s.echo = e

	// Bypass auth middleware by setting user_id and seeding roles in Redis
	e.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "user_id", 1))
			next.ServeHTTP(w, r)
		})
	})

	roles := []string{"Admin_Role_10", "Admin_Admin_14"}
	cache.SetToCache(s.ts.Ctx, cacheStore, "user_roles:1", &roles, 5*time.Minute)

	apiErrorHandler := apierror.NewApiHandler(obs, log)
	roleMapper := response_api.NewRoleResponseMapper()
	gatewayCache := apicache.NewGatewayCache(cacheStore)
	client := pb.NewRoleServiceClient(conn)
	rolehandler.NewHandlerRole(
		e,
		client,
		log,
		roleMapper,
		apiErrorHandler,
		gatewayCache,
	)
}

func (s *RoleApiTestSuite) TearDownSuite() {
	s.conn.Close()
	s.grpcServer.Stop()
	s.redisClient.Close()
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *RoleApiTestSuite) TestRoleApiLifecycle() {
	// 1. Create
	reqBody := requests.CreateRoleRequest{
		Name: "API Role",
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/role", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Require().Equal(http.StatusOK, rec.Code, rec.Body.String())
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	data := res["data"].(map[string]interface{})
	s.Equal(reqBody.Name, data["name"])
	s.roleID = int(data["id"].(float64))

	// 2. FindAll
	req = httptest.NewRequest(http.MethodGet, "/api/role", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 3. FindById
	s.Require().NotZero(s.roleID)
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/role/%d", s.roleID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
	json.Unmarshal(rec.Body.Bytes(), &res)
	data = res["data"].(map[string]interface{})
	s.Equal(float64(s.roleID), data["id"])

	// 4. FindByActive
	req = httptest.NewRequest(http.MethodGet, "/api/role/active", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 5. FindByTrashed
	req = httptest.NewRequest(http.MethodGet, "/api/role/trashed", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 6. FindByUserId
	req = httptest.NewRequest(http.MethodGet, "/api/role/user/1", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 7. Update
	updateBody := requests.UpdateRoleRequest{
		Name: "Updated API Role",
	}
	body, _ = json.Marshal(updateBody)
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/role/update/%d", s.roleID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 8. Restore
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/role/restore/%d", s.roleID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 9. DeletePermanent
	req = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/role/permanent/%d", s.roleID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 10. RestoreAll
	req = httptest.NewRequest(http.MethodPost, "/api/role/restore/all", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 11. DeleteAll
	req = httptest.NewRequest(http.MethodDelete, "/api/role/permanent-all", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func TestRoleApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleApiTestSuite))
}
