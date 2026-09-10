package services

import (
    "context"

    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
)


type SkuServiceInterface interface {
    CreateSKU(ctx context.Context, productId string, OptionIDs []uint, price float64, stock int, token string) (*models.SKU, error)
    UpdateSKU(ctx context.Context, id string, price float64, stock int, token string) (*models.SKU, error)
    GetSKU(ctx context.Context, ID string, token string) (models.SKU, error)
    GetProductSKUList(ctx context.Context, productID string, token string) ([]dto.SKUListItem, error)
}

// BasketService defines sku operations
type SupplierService interface {
  CreateSupplier(ctx context.Context, name, repName, phone, email string) (*models.Supplier, error)
  UpdateSupplier(ctx context.Context, id, name, repName, logoPath, phone string, phoneVerified bool, email string, sms, whatsApp, emailNotify, appPush bool) (*models.Supplier, error)
  GetSupplier(ctx context.Context, id string) (*models.Supplier, error)
  GetSupplierBasic(ctx context.Context, id string) (*dto.SupplierBasic, error)
  GetBasicSupplierList(ctx context.Context, token string) ([]dto.SupplierBasic, error)
  GetMerchantProductSupliers(page, limit, sort string) (*models.Pagination, error)
  DeleteSupplier(ctx context.Context, id string) (error)
}

type ProductService interface {
  CreateVariant(ctx context.Context, id, name string, options []string) (*models.Variant, error)
  GetVariant(ctx context.Context, id string) (*dto.Variant, error)
  GetVariantList(ctx context.Context, id string) ([]dto.VariantGroup, error)
  DeleteVariant(ctx context.Context, id string) (error)

  CreateOptions(ctx context.Context, id string, options []string) (error)
  GetOption(ctx context.Context, optionID string) (*models.Option, error)
  UpdateOption(ctx context.Context, optionID, value string) (models.Option, error)
  DeleteOption(ctx context.Context, optionID string) (error)

  CreateCategory(ctx context.Context, name, alternateKey, iconPath, description string, parentId, subType, sortOrder uint) (*models.Category, error)
  UpdateCategory(ctx context.Context, id, name, alternateKey, iconPath, description string, parentId, subType, sortOrder uint) (*models.Category, error)
  GetCategory(ctx context.Context, id string) (*models.Category, error)
  ListCategories(ctx context.Context) ([]dto.CategoryListItem, error)
  ListProductCategoryItem(ctx context.Context, alternateKey, id string) ([]dto.ProductCategoryItem, error)
  GetCategoryBreadcrumb(ctx context.Context, id string) ([]models.Breadcrumb, error)
  GetMerchantCategoriesByParent(ctx context.Context, id, alt string) ([]dto.CategoryMerchantListItem, error)
  GetCategoriesByParent(ctx context.Context, id, alt string) ([]dto.CategoryListItem, error)
  DeleteCategory(ctx context.Context, id string) error

  ReadProductCategories(ctx context.Context, id string) (dto.CategoriesResponse, error)
  UpdateProductCategories(ctx context.Context, productID uint, categoryIDs []uint) (error)

  ReadProductParameters(ctx context.Context, productID string) ([]dto.ProductParameterItem, error)
  SaveProductParameters(ctx context.Context, productParameters []models.ProductParameter) (error)

  CreateProduct(ctx context.Context, name, description string, supplierID uint) (*models.Product, error)
  ListProducts(ctx context.Context, page, limit, sort, category, token string) (*models.Pagination, error)
  CountProducts(ctx context.Context, category, token string) (*dto.ProductCount, error)
//  ListProducts(ctx context.Context, userId, page, limit, sort, category, token string) (*models.Pagination, error)
  ListMerchantProducts(ctx context.Context, page, limit, sort string) (*models.Pagination, error)
  GetProductBasicInfo(ctx context.Context, id string) (*dto.ProductBasicInfo, error)
  UpdateProduct(ctx context.Context, id, name, description string, state models.ProductState) (*models.Product, error)
  SetState(ctx context.Context, id, state string) (error)

  GetShopProduct(ctx context.Context, productID string) (*dto.ProductShopListItem, error)
  GetShopProductDetail(ctx context.Context, id string) (*dto.ProductResponse, error)
  GetProductFromSKU(ctx context.Context, skuCode, token string) (*dto.ProductVariantDetail,error)
}

type ReviewService interface {
  DeleteReview(ctx context.Context, reviewID string) (error)
  ListReviewsByProduct(ctx context.Context, productID string) ([]dto.ReviewListItem, error)
  ListReviews(ctx context.Context, userID string) ([]dto.ReviewListItem, error)
  GetReview(ctx context.Context, id_str string) (*models.Review, error)
  UpdateReview(ctx context.Context, id_str, comment string, score uint, moderated bool) (*models.Review, error)
  CreateReview(ctx context.Context, score uint, comment, name, accountID string, productID uint) (*models.Review, error)

  ToggleLike(ctx context.Context, productID, accountID string) (bool, error)
  CountLikes(ctx context.Context, productID string) (int64, error)
}

type ProductSearchService interface {
  SearchProducts(ctx context.Context, prefs dto.Preferences, order string) ([]int, error)
  IndexProducts(ctx context.Context) error
}
