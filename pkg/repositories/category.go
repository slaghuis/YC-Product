package repositories

import (
//    "gorm.io/gorm"
    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
)

func (r *ProductRepository) GetCategory(categoryId uint) (*models.Category, error) {
  var item models.Category
  err := r.db.First(&item, categoryId).Error
  return &item, err
}

func (r *ProductRepository) SaveCategory(item *models.Category) error {
  return r.db.Save(item).Error
}

func (r *ProductRepository) GetCategoryBreadcrumb(categoryId uint) ([]models.Breadcrumb, error) {
  // Using raw SQL for recursive CTE
  rawSQL := `
    WITH RECURSIVE category_tree AS (
      SELECT id, name, parent_id FROM categories WHERE id = ?
      UNION ALL
      SELECT c.id, c.name, c.parent_id
      FROM categories c
      JOIN category_tree ct ON c.id = ct.parent_id
      )
      SELECT * FROM category_tree ORDER BY id;
      `
  var breadcrumbs []models.Breadcrumb
  err := r.db.Raw(rawSQL, categoryId).Scan(&breadcrumbs).Error

  return breadcrumbs, err
}

func (r *ProductRepository) DeleteCategory(categoryId uint) (error) {
  return r.db.Delete(&models.Category{}, categoryId).Error
}

func (r *ProductRepository) ListMerchantCategoriesByParentId(parentId uint) ([]dto.CategoryMerchantListItem, error) {
  var items []dto.CategoryMerchantListItem
  err := r.db.Table("categories child").
        Select(`child.alternate_key,
          child.id,
          child.name,
          child.icon_path,
          child.description,
          SUM(likes.score) AS likes_count,
          COUNT(DISTINCT reviews.id) AS reviews_count,
          COUNT(DISTINCT products.id) AS product_count,
          0 AS sales,
          0.0 as earnings
          `).
        Joins("LEFT JOIN product_categories AS pc on pc.category_id = child.id").
        Joins("LEFT JOIN products on products.id = pc.product_id").
        Joins("LEFT JOIN likes on likes.product_id = products.id").
        Joins("LEFT JOIN reviews on reviews.product_id = products.id").
        Where("child.parent_id = ?", parentId).
        Group("child.alternate_key, child.id, child.name, child.icon_path, child.description").
        Order("child.sort_order asc").
        Scan(&items).Error

  return items, err
}

func (r *ProductRepository) ListMerchantCategoriesByAltKey(alt string) ([]dto.CategoryMerchantListItem, error) {
  var items []dto.CategoryMerchantListItem

  err := r.db.Table("categories child").
          Select(`child.alternate_key,
            child.id,
            child.name,
            child.icon_path,
            child.description,
            SUM(likes.score) AS likes_count,
            COUNT(DISTINCT reviews.id) AS reviews_count,
            COUNT(DISTINCT products.id) AS product_count,
            0 AS sales,
            0.0 as earnings
            `).
          Joins("JOIN categories AS parent on child.parent_id = parent.id").
          Joins("LEFT JOIN product_categories AS pc on pc.category_id = child.id").
          Joins("LEFT JOIN products on products.id = pc.product_id").
          Joins("LEFT JOIN likes on likes.product_id = products.id").
          Joins("LEFT JOIN reviews on reviews.product_id = products.id").
          Where("parent.alternate_key = ?", alt).
          Group("child.alternate_key, child.id, child.name, child.icon_path, child.description").
          Order("child.sort_order asc").
          Scan(&items).Error

  return items, err
}

func (r *ProductRepository) ListCategoriesByParentId(parentId uint) ([]dto.CategoryListItem, error) {
  var items []dto.CategoryListItem

  err := r.db.Table("categories child").
        Select("child.alternate_key, child.parent_id, child.name, child.icon_path, child.sub_type, child.description, child.sort_order, count(pc.product_id) as product_count").
        Joins("LEFT JOIN product_categories AS pc on pc.category_id = child.id").
        Joins("LEFT JOIN products AS p on p.id = pc.product_id").
        Where("child.parent_id = ?", parentId).
        Group("child.alternate_key, child.parent_id, child.name, child.icon_path, child.sub_type, child.description, child.sort_order").
        Order("child.sort_order asc").
        Scan(&items).Error
  return items, err
}

func (r *ProductRepository) ListCategoriesByAltKey(alt string) ([]dto.CategoryListItem, error) {
    var items []dto.CategoryListItem
    err := r.db.Table("categories child").
          Select("child.alternate_key, child.parent_id, child.name, child.icon_path, child.sub_type, child.description, child.sort_order, count(pc.product_id) as product_count").
          Joins("JOIN categories AS parent on child.parent_id = parent.id").
          Joins("LEFT JOIN product_categories AS pc on pc.category_id = child.id").
          Joins("LEFT JOIN products AS p on p.id = pc.product_id").
          Where("parent.alternate_key = ?", alt).
          Group("child.alternate_key, child.parent_id, child.name, child.icon_path, child.sub_type, child.description, child.sort_order").
          Order("child.sort_order asc").
          Scan(&items).Error
    return items, err
}

func (r *ProductRepository) ListCategories() ([]dto.CategoryListItem, error) {
    var items []dto.CategoryListItem
    err := r.db.Table("categories").
      Select("alternate_key, parent_id, categories.name, icon_path, sub_type, categories.description, sort_order, count(pc.product_id) as product_count").
      Joins("left join product_categories pc on pc.category_id = categories.id").
      Group("alternate_key, parent_id, categories.name, icon_path, sub_type, categories.description, sort_order").
      Scan(&items).Order("sort_order asc").Error
    return items, err
}

func (r *ProductRepository) ListProductCategoryItem(alternateKey, id string) ([]dto.ProductCategoryItem, error) {
  var items []dto.ProductCategoryItem
  err := r.db.Table("dbo.categories AS parent").
    Select(`
        child.id AS category_id,
        child.name AS category_name,
        child.icon_path,
        CASE WHEN cat.product_id IS NULL THEN FALSE ELSE TRUE END AS selected
    `).
    Joins("INNER JOIN dbo.categories AS child ON child.parent_id = parent.id").
    Joins("LEFT JOIN dbo.product_categories AS cat ON cat.category_id = child.id AND cat.product_id = ?", id).
    Where("parent.alternate_key = ?", alternateKey).
    Order("child.sort_order ASC").
    Scan(&items).Error

  return items, err
}

func (r *ProductRepository) ReadBaseCategories() ([]dto.BaseCategory, error) {
  var baseKeys []dto.BaseCategory
  err := r.db.Table("dbo.categories AS child").
    Select("child.alternate_key, child.name").
    Joins("JOIN dbo.categories AS parent ON parent.id = child.parent_id").
    Where("parent.alternate_key = ?", "base").
    Scan(&baseKeys).Error
  return baseKeys, err
}
