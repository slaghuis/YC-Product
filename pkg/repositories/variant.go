package repositories

import (
    "github.com/slaghuis/YC-Product/pkg/models"
)

func (r *ProductRepository) SaveVariant(item *models.Variant) error {
  return r.db.Save(item).Error
}

func (r *ProductRepository) GetVariant(ID uint) (*models.Variant, error) {
  var item models.Variant
  err := r.db.First(&item, ID).Error
  return &item, err
}

func (r *ProductRepository) GetVariantList(productID string) ([]models.Variant, error) {
  var variants []models.Variant
  err := r.db.Table("variants").
    Select("variants.ID as id, products.id as product_id, variants.name as name, variants.created_at as created_at, variants.updated_at as updated_at").
    Joins("JOIN dbo.products ON products.id = variants.product_id").
    Where("products.id = ?", productID).
    Scan(&variants).Error
  return variants, err
}


func (r *ProductRepository) DeleteVariant(variantID uint) error {
  return r.db.Delete(&models.Variant{}, variantID).Error
}

/*
var item models.Variant
item.ProductID   = productId
item.Name        = body.Name

if result := h.db.Create(&item); result.Error != nil {
    c.AbortWithError(http.StatusInternalServerError, result.Error)
    return
}

var options []models.Option
for _, opt := range body.Options{
  options = append(options, models.Option{
    VariantID: item.ID,
    Value:     opt,
  })
}

if result := h.db.Create(&options); result.Error != nil {
  c.AbortWithError(http.StatusInternalServerError, result.Error)
  return
}

c.JSON(http.StatusCreated, gin.H{
  "message": "Product variant created",
  "variant":	item,
})
*/
