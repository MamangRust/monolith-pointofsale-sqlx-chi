package category_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MamangRust/monolith-pointofsale-apigateway/apierror"
	apicache "github.com/MamangRust/monolith-pointofsale-apigateway/cache/gateway_cache"
	"github.com/MamangRust/monolith-pointofsale-apigateway/handler"
	response_api "github.com/MamangRust/monolith-pointofsale-shared/mapper/response/api"
	pb "github.com/MamangRust/monolith-pointofsale-apigateway/pbcompat"
	tests "github.com/MamangRust/monolith-pointofsale-test"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"
)

type CategoryStatsApiTestSuite struct {
	tests.BaseTestSuite
	echo       chi.Router
	categoryID int
	merchantID int
	userID     int
}

func (s *CategoryStatsApiTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupRoleService()
	s.SetupUserService()
	s.SetupCategoryService()
	s.SetupMerchantService()
	s.SetupProductService()
	s.SetupOrderItemService()
	s.SetupTransactionService()
	s.SetupOrderService()

	s.echo = chi.NewRouter()
	apiHandler := apierror.NewApiHandler(s.Obs, s.Log)

	catMapper := response_api.NewCategoryResponseMapper()
	gatewayCache := apicache.NewGatewayCache(s.GetCacheStore())
	client := pb.NewCategoryServiceClient(s.Conns["category"])
	handler.NewHandlerCategory(
		s.echo,
		client,
		s.Log,
		catMapper,
		apiHandler,
		gatewayCache,
	)
	ctx := context.Background()
	s.userID = s.SeedUser(ctx)
	s.merchantID = s.SeedMerchant(ctx, s.userID)
	s.categoryID = s.SeedCategory(ctx)
	prodID := s.SeedProduct(ctx, s.merchantID, s.categoryID)
	orderID := s.SeedOrder(ctx, s.userID, s.merchantID, prodID)

	// Ensure created_at is set to current time to be picked up by stats
	_, err := s.DBPool().Exec(ctx, "UPDATE orders SET created_at = $1 WHERE order_id = $2",
		time.Now(), orderID)
	s.Require().NoError(err)
}

func (s *CategoryStatsApiTestSuite) TestFindMonthTotalPrice() {
	now := time.Now()
	url := fmt.Sprintf("/api/category/monthly-total-pricing?year=%d&month=%d", now.Year(), int(now.Month()))
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])
	s.NotEmpty(res["data"])
}

func (s *CategoryStatsApiTestSuite) TestFindYearTotalPrice() {
	year := time.Now().Year()
	url := fmt.Sprintf("/api/category/yearly-total-pricing?year=%d", year)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])
	s.NotEmpty(res["data"])
}

func (s *CategoryStatsApiTestSuite) TestFindMonthPrice() {
	year := time.Now().Year()
	url := fmt.Sprintf("/api/category/monthly-pricing?year=%d", year)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])
	s.NotEmpty(res["data"])
}

func (s *CategoryStatsApiTestSuite) TestFindYearPrice() {
	year := time.Now().Year()
	url := fmt.Sprintf("/api/category/yearly-pricing?year=%d", year)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])
	s.NotEmpty(res["data"])
}

func (s *CategoryStatsApiTestSuite) TestFindMonthTotalPriceById() {
	now := time.Now()
	url := fmt.Sprintf("/api/category/mycategory/monthly-total-pricing?year=%d&month=%d&category_id=%d",
		now.Year(), int(now.Month()), s.categoryID)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])
	s.NotEmpty(res["data"])
}

func (s *CategoryStatsApiTestSuite) TestFindMonthTotalPriceByMerchant() {
	now := time.Now()
	url := fmt.Sprintf("/api/category/merchant/monthly-total-pricing?year=%d&month=%d&merchant_id=%d",
		now.Year(), int(now.Month()), s.merchantID)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])
	s.NotEmpty(res["data"])
}

func TestCategoryStatsApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CategoryStatsApiTestSuite))
}
