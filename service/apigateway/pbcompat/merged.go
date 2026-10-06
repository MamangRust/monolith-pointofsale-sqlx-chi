package pbcompat

// Unified service clients: the flat shared/pb layout exposed one client per
// entity, while the local pb module splits them into Command/Query/Stats
// services. These merged interfaces keep handler code unchanged.
import (
	"google.golang.org/grpc"

	rootpb "github.com/MamangRust/monolith-pointofsale-pb"
	cashierpb "github.com/MamangRust/monolith-pointofsale-pb/cashier"
	categorypb "github.com/MamangRust/monolith-pointofsale-pb/category"
	merchantpb "github.com/MamangRust/monolith-pointofsale-pb/merchant"
	merchant_documentpb "github.com/MamangRust/monolith-pointofsale-pb/merchant_document"
	orderpb "github.com/MamangRust/monolith-pointofsale-pb/order"
	productpb "github.com/MamangRust/monolith-pointofsale-pb/product"
	rolepb "github.com/MamangRust/monolith-pointofsale-pb/role"
	transactionpb "github.com/MamangRust/monolith-pointofsale-pb/transaction"
	userpb "github.com/MamangRust/monolith-pointofsale-pb/user"
)

// CashierServiceClient
type CashierServiceClient interface {
	cashierpb.CashierCommandServiceClient
	cashierpb.CashierQueryServiceClient
	cashierpb.CashierStatsByIdServiceClient
	cashierpb.CashierStatsByMerchantServiceClient
	cashierpb.CashierStatsServiceClient
}

type cashierCombined struct {
	cashierpb.CashierCommandServiceClient
	cashierpb.CashierQueryServiceClient
	cashierpb.CashierStatsByIdServiceClient
	cashierpb.CashierStatsByMerchantServiceClient
	cashierpb.CashierStatsServiceClient
}

func NewCashierServiceClient(cc grpc.ClientConnInterface) CashierServiceClient {
	return cashierCombined{
		CashierCommandServiceClient:         cashierpb.NewCashierCommandServiceClient(cc),
		CashierQueryServiceClient:           cashierpb.NewCashierQueryServiceClient(cc),
		CashierStatsByIdServiceClient:       cashierpb.NewCashierStatsByIdServiceClient(cc),
		CashierStatsByMerchantServiceClient: cashierpb.NewCashierStatsByMerchantServiceClient(cc),
		CashierStatsServiceClient:           cashierpb.NewCashierStatsServiceClient(cc),
	}
}

// CategoryServiceClient
type CategoryServiceClient interface {
	categorypb.CategoryCommandServiceClient
	categorypb.CategoryQueryServiceClient
	categorypb.CategoryStatsByIdServiceClient
	categorypb.CategoryStatsByMerchantServiceClient
	categorypb.CategoryStatsServiceClient
}

type categoryCombined struct {
	categorypb.CategoryCommandServiceClient
	categorypb.CategoryQueryServiceClient
	categorypb.CategoryStatsByIdServiceClient
	categorypb.CategoryStatsByMerchantServiceClient
	categorypb.CategoryStatsServiceClient
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return categoryCombined{
		CategoryCommandServiceClient:         categorypb.NewCategoryCommandServiceClient(cc),
		CategoryQueryServiceClient:           categorypb.NewCategoryQueryServiceClient(cc),
		CategoryStatsByIdServiceClient:       categorypb.NewCategoryStatsByIdServiceClient(cc),
		CategoryStatsByMerchantServiceClient: categorypb.NewCategoryStatsByMerchantServiceClient(cc),
		CategoryStatsServiceClient:           categorypb.NewCategoryStatsServiceClient(cc),
	}
}

// MerchantServiceClient
type MerchantServiceClient interface {
	merchantpb.MerchantCommandServiceClient
	merchantpb.MerchantQueryServiceClient
}

type merchantCombined struct {
	merchantpb.MerchantCommandServiceClient
	merchantpb.MerchantQueryServiceClient
}

func NewMerchantServiceClient(cc grpc.ClientConnInterface) MerchantServiceClient {
	return merchantCombined{
		MerchantCommandServiceClient: merchantpb.NewMerchantCommandServiceClient(cc),
		MerchantQueryServiceClient:   merchantpb.NewMerchantQueryServiceClient(cc),
	}
}

// MerchantDocumentServiceClient
type MerchantDocumentServiceClient interface {
	merchant_documentpb.MerchantDocumentCommandServiceClient
	merchant_documentpb.MerchantDocumentQueryServiceClient
}

type merchantDocumentCombined struct {
	merchant_documentpb.MerchantDocumentCommandServiceClient
	merchant_documentpb.MerchantDocumentQueryServiceClient
}

func NewMerchantDocumentServiceClient(cc grpc.ClientConnInterface) MerchantDocumentServiceClient {
	return merchantDocumentCombined{
		MerchantDocumentCommandServiceClient: merchant_documentpb.NewMerchantDocumentCommandServiceClient(cc),
		MerchantDocumentQueryServiceClient:   merchant_documentpb.NewMerchantDocumentQueryServiceClient(cc),
	}
}

// OrderServiceClient
type OrderServiceClient interface {
	orderpb.OrderCommandServiceClient
	orderpb.OrderQueryServiceClient
	orderpb.OrderStatsByMerchantServiceClient
	orderpb.OrderStatsServiceClient
}

type orderCombined struct {
	orderpb.OrderCommandServiceClient
	orderpb.OrderQueryServiceClient
	orderpb.OrderStatsByMerchantServiceClient
	orderpb.OrderStatsServiceClient
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return orderCombined{
		OrderCommandServiceClient:         orderpb.NewOrderCommandServiceClient(cc),
		OrderQueryServiceClient:           orderpb.NewOrderQueryServiceClient(cc),
		OrderStatsByMerchantServiceClient: orderpb.NewOrderStatsByMerchantServiceClient(cc),
		OrderStatsServiceClient:           orderpb.NewOrderStatsServiceClient(cc),
	}
}

// OrderItemServiceClient
type OrderItemServiceClient = rootpb.OrderItemQueryServiceClient

var NewOrderItemServiceClient = rootpb.NewOrderItemQueryServiceClient

// ProductServiceClient
type ProductServiceClient interface {
	productpb.ProductCommandServiceClient
	productpb.ProductQueryServiceClient
}

type productCombined struct {
	productpb.ProductCommandServiceClient
	productpb.ProductQueryServiceClient
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return productCombined{
		ProductCommandServiceClient: productpb.NewProductCommandServiceClient(cc),
		ProductQueryServiceClient:   productpb.NewProductQueryServiceClient(cc),
	}
}

// RoleServiceClient
type RoleServiceClient interface {
	rolepb.RoleCommandServiceClient
	rolepb.RoleQueryServiceClient
}

type roleCombined struct {
	rolepb.RoleCommandServiceClient
	rolepb.RoleQueryServiceClient
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return roleCombined{
		RoleCommandServiceClient: rolepb.NewRoleCommandServiceClient(cc),
		RoleQueryServiceClient:   rolepb.NewRoleQueryServiceClient(cc),
	}
}

// TransactionServiceClient
type TransactionServiceClient interface {
	transactionpb.TransactionCommandServiceClient
	transactionpb.TransactionQueryServiceClient
	transactionpb.TransactionStatsByMerchantServiceClient
	transactionpb.TransactionStatsServiceClient
}

type transactionCombined struct {
	transactionpb.TransactionCommandServiceClient
	transactionpb.TransactionQueryServiceClient
	transactionpb.TransactionStatsByMerchantServiceClient
	transactionpb.TransactionStatsServiceClient
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return transactionCombined{
		TransactionCommandServiceClient:         transactionpb.NewTransactionCommandServiceClient(cc),
		TransactionQueryServiceClient:           transactionpb.NewTransactionQueryServiceClient(cc),
		TransactionStatsByMerchantServiceClient: transactionpb.NewTransactionStatsByMerchantServiceClient(cc),
		TransactionStatsServiceClient:           transactionpb.NewTransactionStatsServiceClient(cc),
	}
}

// UserServiceClient
type UserServiceClient interface {
	userpb.UserCommandServiceClient
	userpb.UserQueryServiceClient
}

type userCombined struct {
	userpb.UserCommandServiceClient
	userpb.UserQueryServiceClient
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return userCombined{
		UserCommandServiceClient: userpb.NewUserCommandServiceClient(cc),
		UserQueryServiceClient:   userpb.NewUserQueryServiceClient(cc),
	}
}
