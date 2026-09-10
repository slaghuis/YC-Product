package repositories

import (
    "gorm.io/gorm"
    "gorm.io/gorm/clause"

    "github.com/lib/pq"

    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
)

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProduct(productID uint) (*models.Product, error) {
  var product models.Product
  err := r.db.First(&product, productID).Error
  return &product, err
}

//Specific funcion to return all products that must be indexed
/*
func (r *ProductRepository) GetSearchProductList() ([]dto.ProductListItem, error) {
  var productList []dto.ProductListItem
  err := r.db.Model(&models.Product{}).
    Select("DISTINCT products.id, products.name, products.description, products.state").
    Where("products.state >= ?", models.ProductStateActive).
    Joins("LEFT JOIN product_categories AS pc on pc.product_id = products.id").
    Joins("LEFT JOIN categories on categories.id = pc.category_id").
    Scan(&productList).Error

  return productList, err
}
*/


type ProductIndexProjection struct {
    ProductID    uint
    ProductName  string
    Supplier     string
    Promoted     bool
    Age          uint
    Categories   pq.StringArray `gorm:"type:text[]"`
    OptionValues pq.StringArray `gorm:"type:text[]"`
}

func (r *ProductRepository) GetProductIndexProjection() ([]ProductIndexProjection, error) {
    var results []ProductIndexProjection

    err := r.db.
        Table("dbo.products").
        Select(`
            products.id as product_id,
            products.name as product_name,
            (products.state = 400) as promoted,
            CURRENT_DATE - products.created_at::date as age,
            suppliers.name as supplier,
            COALESCE(array_agg(DISTINCT categories.alternate_key) FILTER (WHERE categories.alternate_key IS NOT NULL), ARRAY[]::text[]) as categories,
            COALESCE(array_agg(DISTINCT options.value) FILTER (WHERE options.value IS NOT NULL), ARRAY[]::text[]) as option_values
        `).
        Joins("LEFT JOIN dbo.product_categories pc ON pc.product_id = products.id").
        Joins("LEFT JOIN dbo.categories ON categories.id = pc.category_id").
        Joins("LEFT JOIN dbo.variants ON variants.product_id = products.id").
        Joins("LEFT JOIN dbo.options ON options.variant_id = variants.id").
        Joins("LEFT JOIN dbo.suppliers ON products.supplier_id=suppliers.id").
        Where("products.state >= ?", models.ProductStateActive).
        Group("products.id, products.name, products.state, suppliers.name").
        Scan(&results).Error

    return results, err
}

// UpdateProductCategories updates the categories for a given product.
// It inserts new associations (upsert) and deletes removed ones in a transaction.
func (r *ProductRepository) UpdateProductCategories( productID uint, categoryIDs []uint) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        // Step 1: Insert new associations (upsert)
        newCategories := make([]dto.ProductCategory, len(categoryIDs))
        for i, cid := range categoryIDs {
            newCategories[i] = dto.ProductCategory{ProductID: productID, CategoryID: cid}
        }

        if err := tx.Clauses(clause.OnConflict{
            Columns:   []clause.Column{{Name: "product_id"}, {Name: "category_id"}},
            DoNothing: true,
        }).Create(&newCategories).Error; err != nil {
            return err
        }

        // Step 2: Delete associations not in the new set
        if err := tx.Where("product_id = ? AND category_id NOT IN ?", productID, categoryIDs).
            Delete(&dto.ProductCategory{}).Error; err != nil {
            return err
        }

        return nil
    })
}


func (r *ProductRepository) SaveProduct(item *models.Product) error {
  return r.db.Save(item).Error
}

func (r *ProductRepository) GetProductList(pagination *models.Pagination, category string) ([]dto.ProductListItem, error) {
  var productList []dto.ProductListItem
  err := r.db.Scopes(models.Paginate(&models.Product{}, pagination),
                    models.FilterByCategoryID(category)).
    Select("DISTINCT products.id, products.name, products.description, products.state").
    Where("products.state >= ?", models.ProductStateActive).
    Joins("LEFT JOIN product_categories AS pc on pc.product_id = products.id").
    Joins("LEFT JOIN categories on categories.id = pc.category_id").
    Scan(&productList).Error

  return productList, err
}

func (r *ProductRepository) GetProductListFromSlice(pagination *models.Pagination, productIDs []int) ([]dto.ProductListItem, error) {
  var productList []dto.ProductListItem
  //err := r.db.Scopes(models.Paginate(&models.Product{}, pagination)).
  err := r.db.
      Table("dbo.products").
    Select("DISTINCT products.id, products.name, products.description, products.state").
    Where("products.id in ?", productIDs).
    Joins("LEFT JOIN product_categories AS pc on pc.product_id = products.id").
    Joins("LEFT JOIN categories on categories.id = pc.category_id").
    Scan(&productList).Error

  return productList, err
}

func (r *ProductRepository) GetProductAndCategories(productID int) (dto.ProductListItem, error) {
  var product dto.ProductListItem
  //err := r.db.Scopes(models.Paginate(&models.Product{}, pagination)).
  err := r.db.
      Table("dbo.products").
    Select("products.id, products.name, products.description, products.state").
    Where("products.id = ?", productID).
    Joins("LEFT JOIN product_categories AS pc on pc.product_id = products.id").
    Joins("LEFT JOIN categories on categories.id = pc.category_id").
    Limit(1).
    Scan(&product).Error

  return product, err
}

func (r *ProductRepository) GetMerchantProdictList(pagination *models.Pagination) ([]dto.ProductMerchantListItem, error) {
  var productList []dto.ProductMerchantListItem
  err := r.db.Scopes(models.Paginate(&models.Product{}, pagination)).
    Select(`products.id,
      products.name,
      products.description,
      products.state,
      SUM(likes.score) AS likes_count,
      COUNT(DISTINCT reviews.id) AS reviews_count,
      COALESCE(images.image_path, '') AS image_path,
      COALESCE(MIN(skus.price),0) AS min_price,
      COALESCE(MAX(skus.price),0) AS max_price,
      0 AS sales,
      0 AS earnings`).
    Joins("LEFT JOIN dbo.images on images.product_id = products.id  AND images.type = 200").
    Joins("LEFT JOIN dbo.likes on likes.product_id = products.id").
    Joins("LEFT JOIN dbo.reviews on reviews.product_id = products.id").
    Joins("LEFT JOIN dbo.skus on skus.product_id = products.id").
    Group("products.id, products.name, products.description, products.state, images.image_path").
    Scan(&productList).Error

  return productList, err
}

func (r *ProductRepository) SetState(productID uint, newState int) (error) {
  return r.db.Model(&models.Product{}).Where("id = ?", productID).Update("state", newState).Error
}

func (r *ProductRepository) GetProductBasicInfo(productID string) (*dto.ProductBasicInfo, error) {
  var definition dto.ProductBasicInfo

  err := r.db.Table("dbo.products").
  Select(`
    products.id AS id,
    products.name AS name,
    products.description AS description,
    products.state AS state,
    products.supplier_id AS supplierid,
    suppliers.name AS suppliername,
    COALESCE(SUM(likes.score), 0) AS likecount,
    COUNT(DISTINCT reviews.id) AS reviewcount
    `).
    Joins("LEFT JOIN dbo.suppliers ON suppliers.id = products.supplier_id").
    Joins("LEFT JOIN dbo.likes ON likes.product_id = products.id").
    Joins("LEFT JOIN dbo.reviews ON reviews.product_id = products.id").
    Where("products.id = ?", productID).
    Group("products.id, products.name, products.description, products.state, suppliers.id, suppliers.name").
    Limit(1).
    Scan(&definition).Error

  return &definition, err
}


type ProductScanRow struct {
    ID           uint             `gorm:"column:id"`
    Name         string           `gorm:"column:name"`
    Description  string           `gorm:"column:description"`
    Promoted     bool             `gorm:"column:promoted"`
    ImagePath    string           `gorm:"column:image_path"`
    MinPrice     float64          `gorm:"column:min_price"`
    MaxPrice     float64          `gorm:"column:max_price"`
    LikesCount   int              `gorm:"column:likes_count"`
    ReviewsCount int              `gorm:"column:reviews_count"`
    SKUCount     int              `gorm:"column:sku_count"`
    SKUsJSON     string           `gorm:"column:skus"`
}

func (r *ProductRepository) GetShopProduct(productID string) (*ProductScanRow, error) {
  var row ProductScanRow

  err := r.db.Raw(`
   SELECT
      p.id,
      p.name,
      p.description,
      (p.state = 400) AS promoted,
      COALESCE(img.image_path, '') AS image_path,
      COALESCE(sku_agg.min_price, 0) AS min_price,
      COALESCE(sku_agg.max_price, 0) AS max_price,
      COALESCE(like_agg.total_score, 0) AS likes_count,
      COALESCE(review_agg.cnt, 0) AS reviews_count,
      COALESCE(sku_agg.sku_count, 0) AS sku_count,
      COALESCE(sku_agg.skus_json, '[]'::json) AS skus

   FROM products p

   LEFT JOIN LATERAL (
     SELECT image_path
     FROM dbo.images
     WHERE product_id = p.id and type = 200
     LIMIT 1
   ) img ON true

   LEFT JOIN LATERAL (
    SELECT
        COUNT(l.id) AS cnt,
        COALESCE(SUM(l.score), 0) AS total_score
    FROM likes l
    WHERE l.product_id = p.id
   ) like_agg ON true

   LEFT JOIN LATERAL (
    SELECT COUNT(r.id) AS cnt
    FROM reviews r
    WHERE r.product_id = p.id
   ) review_agg ON true

   JOIN LATERAL (
    SELECT
        MIN(s.price) AS min_price,
        MAX(s.price) AS max_price,
        COUNT(s.id) AS sku_count,
        json_agg(
            jsonb_build_object(
                'id', s.id,
                'code', s.code,
                'price', s.price,
                'stock', s.stock,
                'options', COALESCE(
                    (SELECT jsonb_agg(
                        jsonb_build_object(
                            'name', v.name,
                            'value', o.value
                        )
                    )
                    FROM sku_options so
                    JOIN options o ON o.id = so.option_id
                    JOIN variants v ON v.id = o.variant_id
                    WHERE so.sk_uid = s.id
                    ), '[]'::jsonb)
            )
        ) AS skus_json
    FROM skus s
    WHERE s.product_id = p.id
    HAVING COUNT(s.id) > 0
   ) sku_agg ON true

   WHERE p.id = $1
   LIMIT 1
  `, productID).Scan(&row).Error

  return &row, err
}


func (r *ProductRepository)GetShopProductDetail(productID uint) (*models.Product, error) {
  var product models.Product
  if err := r.db.Preload("Variants.Options").
                 Preload("Supplier").
                 Preload("ProductParameter.Category").
                 Preload("Categories", "parent_id IN (?)",
                  r.db.Model(&models.Category{}).
                  Select("id").
                  Where("alternate_key = ?", "lifestyle")).
                 First(&product, productID).Error; err != nil {
    return nil, err
  }

  return &product, nil
}
