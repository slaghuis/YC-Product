package repositories

import (
    "gorm.io/gorm"
    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
)

type SupplierRepository struct {
    db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
    return &SupplierRepository{db: db}
}

func (r *SupplierRepository) SaveSupplier(item *models.Supplier) error {
  return r.db.Save(item).Error
}

func (r *SupplierRepository) GetSupplier(id uint) (*models.Supplier, error) {
  var item models.Supplier
  err := r.db.First(&item, id).Error
  return &item, err
}

func (r *SupplierRepository) GetSupplierBasic(id uint) (*dto.SupplierBasic, error) {
  var item dto.SupplierBasic
  err := r.db.Select("id", "name", "rep_name", "logo_path", "phone", "phone_verified", "email", "sms", "whats_app", "email_notify", "app_push").
      First(&item, id).Error
  return &item, err
}

func (r *SupplierRepository) GetSupplierBasicList() ([]dto.SupplierBasic, error) {
    var supplierList []dto.SupplierBasic
    err := r.db.Table("dbo.suppliers").
      Select(`suppliers.id,
        suppliers.name,
        suppliers.rep_name,
        suppliers.logo_path,
        suppliers.phone,
        suppliers.phone_verified,
        suppliers.email,
        suppliers.sms,
        suppliers.whats_app,
        suppliers.email_notify,
        suppliers.app_push,
        COUNT(DISTINCT products.id) AS product_count`).
      Joins("LEFT JOIN products on suppliers.id = products.supplier_id").
      Group(`suppliers.id,
        suppliers.name,
        suppliers.rep_name,
        suppliers.logo_path,
        suppliers.phone,
        suppliers.phone_verified,
        suppliers.email,
        suppliers.sms,
        suppliers.whats_app,
        suppliers.email_notify,
        suppliers.app_push`).
      Order("suppliers.name asc").
      Scan(&supplierList).Error
    return supplierList, err
}

func (r *SupplierRepository) DeleteSupplier(id uint) (error) {
  return r.db.Delete(&models.Supplier{}, id).Error
}

func (r *SupplierRepository) GetMerchantProductSupliers(pagination *models.Pagination) ([]dto.SupplierMerchantListItem, error) {
  var supplierList []dto.SupplierMerchantListItem
  err := r.db.Scopes(models.Paginate(&models.Supplier{}, pagination)).
    Select(`suppliers.id,
      suppliers.name,
      suppliers.rep_name,
      suppliers.logo_path,
      SUM(likes.score) AS likes_count,
      COUNT(DISTINCT reviews.id) AS reviews_count,
      COUNT(DISTINCT products.id) AS product_count,
      0 AS sales,
      0.0 as earnings
      `).
    Joins("LEFT JOIN products on suppliers.id = products.supplier_id").
    Joins("LEFT JOIN likes on likes.product_id = products.id").
    Joins("LEFT JOIN reviews on reviews.product_id = products.id").
    Group("suppliers.id, suppliers.name, suppliers.rep_name, suppliers.logo_path").
    Scan(&supplierList).Error

  return supplierList, err
}
