package pbcompat

// Package pbcompat re-exports the split per-service pb packages under the
// flat names the handlers were written against (former shared/pb layout),
// so handlers can keep referencing every type as pb.X.
import (
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

// rootpb
type AuthServiceClient = rootpb.AuthServiceClient
type ForgotPasswordRequest = rootpb.ForgotPasswordRequest
type GetMeRequest = rootpb.GetMeRequest
type LoginRequest = rootpb.LoginRequest
type RefreshTokenRequest = rootpb.RefreshTokenRequest
type RegisterRequest = rootpb.RegisterRequest
type ResetPasswordRequest = rootpb.ResetPasswordRequest
type VerifyCodeRequest = rootpb.VerifyCodeRequest

var NewAuthServiceClient = rootpb.NewAuthServiceClient

// cashierpb
type CreateCashierRequest = cashierpb.CreateCashierRequest
type FindAllCashierRequest = cashierpb.FindAllCashierRequest
type FindByIdCashierRequest = cashierpb.FindByIdCashierRequest
type FindYearCashier = cashierpb.FindYearCashier
type FindYearCashierById = cashierpb.FindYearCashierById
type FindYearCashierByMerchant = cashierpb.FindYearCashierByMerchant
type FindYearMonthTotalSales = cashierpb.FindYearMonthTotalSales
type FindYearMonthTotalSalesById = cashierpb.FindYearMonthTotalSalesById
type FindYearMonthTotalSalesByMerchant = cashierpb.FindYearMonthTotalSalesByMerchant
type FindYearTotalSales = cashierpb.FindYearTotalSales
type FindYearTotalSalesById = cashierpb.FindYearTotalSalesById
type FindYearTotalSalesByMerchant = cashierpb.FindYearTotalSalesByMerchant
type UpdateCashierRequest = cashierpb.UpdateCashierRequest

// categorypb
type CreateCategoryRequest = categorypb.CreateCategoryRequest
type FindAllCategoryRequest = categorypb.FindAllCategoryRequest
type FindByIdCategoryRequest = categorypb.FindByIdCategoryRequest
type FindYearCategory = categorypb.FindYearCategory
type FindYearCategoryById = categorypb.FindYearCategoryById
type FindYearCategoryByMerchant = categorypb.FindYearCategoryByMerchant
type FindYearMonthTotalPrices = categorypb.FindYearMonthTotalPrices
type FindYearMonthTotalPriceById = categorypb.FindYearMonthTotalPriceById
type FindYearMonthTotalPriceByMerchant = categorypb.FindYearMonthTotalPriceByMerchant
type FindYearTotalPrices = categorypb.FindYearTotalPrices
type FindYearTotalPriceById = categorypb.FindYearTotalPriceById
type FindYearTotalPriceByMerchant = categorypb.FindYearTotalPriceByMerchant
type UpdateCategoryRequest = categorypb.UpdateCategoryRequest

// merchantpb
type CreateMerchantRequest = merchantpb.CreateMerchantRequest
type FindAllMerchantRequest = merchantpb.FindAllMerchantRequest
type FindByIdMerchantRequest = merchantpb.FindByIdMerchantRequest
type UpdateMerchantRequest = merchantpb.UpdateMerchantRequest
type UpdateMerchantStatusRequest = merchantpb.UpdateMerchantStatusRequest

// merchant_documentpb
type CreateMerchantDocumentRequest = merchant_documentpb.CreateMerchantDocumentRequest
type DeleteMerchantDocumentPermanentRequest = merchant_documentpb.DeleteMerchantDocumentPermanentRequest
type FindAllMerchantDocumentsRequest = merchant_documentpb.FindAllMerchantDocumentsRequest
type FindMerchantDocumentByIdRequest = merchant_documentpb.FindMerchantDocumentByIdRequest
type RestoreMerchantDocumentRequest = merchant_documentpb.RestoreMerchantDocumentRequest
type TrashedMerchantDocumentRequest = merchant_documentpb.TrashedMerchantDocumentRequest
type UpdateMerchantDocumentRequest = merchant_documentpb.UpdateMerchantDocumentRequest
type UpdateMerchantDocumentStatusRequest = merchant_documentpb.UpdateMerchantDocumentStatusRequest

// order_itempb (root package)
type FindAllOrderItemRequest = rootpb.FindAllOrderItemRequest
type FindByIdOrderItemRequest = rootpb.FindByIdOrderItemRequest

// orderpb
type CreateOrderRequest = orderpb.CreateOrderRequest
type CreateOrderItemRequest = orderpb.CreateOrderItemRequest
type UpdateOrderItemRequest = orderpb.UpdateOrderItemRequest
type FindAllOrderRequest = orderpb.FindAllOrderRequest
type FindByIdOrderRequest = orderpb.FindByIdOrderRequest
type FindYearOrder = orderpb.FindYearOrder
type FindYearOrderByMerchant = orderpb.FindYearOrderByMerchant
type FindYearMonthTotalRevenue = orderpb.FindYearMonthTotalRevenue
type FindYearMonthTotalRevenueByMerchant = orderpb.FindYearMonthTotalRevenueByMerchant
type FindYearTotalRevenue = orderpb.FindYearTotalRevenue
type FindYearTotalRevenueByMerchant = orderpb.FindYearTotalRevenueByMerchant
type UpdateOrderRequest = orderpb.UpdateOrderRequest

// productpb
type CreateProductRequest = productpb.CreateProductRequest
type FindAllProductRequest = productpb.FindAllProductRequest
type FindAllProductCategoryRequest = productpb.FindAllProductCategoryRequest
type FindAllProductMerchantRequest = productpb.FindAllProductMerchantRequest
type FindByIdProductRequest = productpb.FindByIdProductRequest
type UpdateProductRequest = productpb.UpdateProductRequest

// rolepb
type CreateRoleRequest = rolepb.CreateRoleRequest
type FindAllRoleRequest = rolepb.FindAllRoleRequest
type FindByIdRoleRequest = rolepb.FindByIdRoleRequest
type FindByIdUserRoleRequest = rolepb.FindByIdUserRoleRequest
type UpdateRoleRequest = rolepb.UpdateRoleRequest

// transactionpb
type CreateTransactionRequest = transactionpb.CreateTransactionRequest
type FindAllTransactionRequest = transactionpb.FindAllTransactionRequest
type FindAllTransactionMerchantRequest = transactionpb.FindAllTransactionMerchantRequest
type FindByIdTransactionRequest = transactionpb.FindByIdTransactionRequest
type FindMonthlyTransactionStatus = transactionpb.FindMonthlyTransactionStatus
type FindMonthlyTransactionStatusByMerchant = transactionpb.FindMonthlyTransactionStatusByMerchant
type FindYearlyTransactionStatus = transactionpb.FindYearlyTransactionStatus
type FindYearlyTransactionStatusByMerchant = transactionpb.FindYearlyTransactionStatusByMerchant
type MonthTransactionMethod = transactionpb.MonthTransactionMethod
type MonthTransactionMethodByMerchant = transactionpb.MonthTransactionMethodByMerchant
type YearTransactionMethod = transactionpb.YearTransactionMethod
type YearTransactionMethodByMerchant = transactionpb.YearTransactionMethodByMerchant
type UpdateTransactionRequest = transactionpb.UpdateTransactionRequest

// userpb
type CreateUserRequest = userpb.CreateUserRequest
type FindAllUserRequest = userpb.FindAllUserRequest
type FindByIdUserRequest = userpb.FindByIdUserRequest
type UpdateUserRequest = userpb.UpdateUserRequest
