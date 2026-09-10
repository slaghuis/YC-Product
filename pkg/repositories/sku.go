package repositories

import (
    "gorm.io/gorm"
    "github.com/slaghuis/YC-Product/pkg/models"
)

type SkuRepository struct {
    db *gorm.DB
}

func NewSkuRepository(db *gorm.DB) *SkuRepository {
    return &SkuRepository{db: db}
}

/*
func (r *SkuRepository) FindByID(ID string) (*models.SKU, error) {
  var item models.SKU
  err := r.db.First(&item, ID).Error
  return &item, err
}
*/

func (r *SkuRepository) FindByID(id string) (*models.SKU, error) {
  var sku models.SKU

  err := r.db.Preload("Options.Option.Variant").
    Preload("Product").
    First(&sku, "id = ?", id).Error
  return &sku, err
}

func (r *SkuRepository) FindByCode(skuCode string) (*models.SKU, error) {
  var sku models.SKU

  err := r.db.Preload("Options.Option.Variant").
    Preload("Product").
    First(&sku, "code = ?", skuCode).Error
  return &sku, err
}


func (r *SkuRepository) SaveSKU(item *models.SKU) error {
  return r.db.Save(item).Error
}

func (r *SkuRepository) SaveSKUOption(item *models.SKUOption) error {
  return r.db.Save(item).Error
}

func (r *SkuRepository) GetProductSKUList(id string) ([]models.SKU, error) {
  var skuList []models.SKU
  err := r.db.Table("dbo.skus").
  Select(`skus.id as id,
    skus.product_id,
    skus.code,
    skus.price,
    skus.stock`).
  Joins("JOIN dbo.products ON skus.product_id=products.id").
  Where("products.id = ?", id).
  Scan(&skuList).Error

  return skuList, err
}
