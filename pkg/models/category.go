package models

import "time"

type Category struct{
  ID                uint                 `gorm:"primaryKey";json:"id"`
  AlternateKey      string               `gorm:"default:null";json:"alternate_key"`
  ParentId          uint                 `gorm:"default:null";json:"parent_id"`
  Name              string               `gorm:"not null;default:null";json:"name"`
  IconPath          string               `gorm:"default:null";json:"icon_path"`
  SubType           uint                 `gorm:"not null";json:"sub_type"`
  Description       string               `json:"description"`
  SortOrder         uint                 `gorm:default:1;json:"sort_order"`
  Products          []Product            `gorm:"many2many:product_categories;";json:"products"`
  ProductParameter  []ProductParameter   `gorm:"foreignKey:CategoryID";json:"parameters"`

//  Children         []Category          `gorm:"foreignKey:ParentId" json:"children"`

  CreatedAt         time.Time            `json:"created_at"`
  UpdatedAt         time.Time            `json:"updated_at"`
}

type Breadcrumb struct {
  ID          uint
  Name        string
  ParentId    uint
}

/*
type CategoryList struct {
    CategoryID  uint   `gorm:\"column:category_id\"`
    DisplayName string `gorm:\"column:display_name\"`
}

type ProductCategoryList struct {
    ParameterName string
    CategoryID    uint
    ID            uint
    Markdown      string
}

type Breadcrumb struct {
  ID          uint
  Name        string
  ParentId    uint
}

// Start of CRUD interface
func GetCategories(db *gorm.DB, item *[]Category) (err error) {
	err = db.Find(item).Error
	if err != nil {
		return err
	}
	return nil
}

// get account by id
func GetCategory(db *gorm.DB, item *Category, id string) (err error) {
	err = db.Where("id = ?", id).First(item).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(db *gorm.DB, item *Category) (err error) {
	err = db.Save(item).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateCategory(db *gorm.DB, item *Category) (*Category, error) {
	err := db.Create(&item).Error
	if err != nil {
		return &Category{}, err
	}
	return item, nil
}

func DeleteCategory(db *gorm.DB, item *Category) (err error) {
	err = db.Delete(item).Error
	if err != nil {
		return err
	}
	return nil
}

// end of CRUD iterface

// Start of utility functions

func GetCategoriesBySubType(db *gorm.DB, item *[]Category, SubType uint) (err error) {
	err = db.Where("sub_type = ?", SubType).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategoryIdFromAlternateKey(db *gorm.DB, alternateKey string) (uint, error) {
  var itemID uint
  err := db.Table("categories").Select("id").Where("alternate_key = ?", alternateKey).Take(&itemID).Error
  if err != nil {
    return 0, err
  }
  return itemID, err
}

func GetCategoriesByParentAlternateKey(db *gorm.DB, AlternateKey string) ( []Category, error) {
	var results []Category
	err := db.Table("categories parent").
         Select("child.*").
    	   Joins("inner join categories child on child.parent_id = parent.id").
    	   Where("parent.alternate_key = ?", AlternateKey).
    	   Scan(&results).Error
	return results, err
}

func GetCategoriesByParentId(db *gorm.DB, parentId string) ( []Category, error) {
	var results []Category
  err := db.Where("parent_id = ?", parentId).Find(&results).Error
  return results, err
}

func GetCategoryListByParentAlternateKey(db *gorm.DB, AlternateKey string) ( []CategoryList, error) {
	var results []CategoryList
	err := db.Table("categories parent").
    		 Select("child.id as category_id, child.name as display_name").
    	   Joins("inner join categories child on child.parent_id = parent.id").
    	   Where("parent.alternate_key = ?", AlternateKey).
    	   Scan(&results).Error
	return results, err
}

func GetCategoryBreadcrumb(db * gorm.DB, categoryId string) ([]Breadcrumb, error) {
  var breadcrumbs []Breadcrumb

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

  result := db.Raw(rawSQL, categoryId).Scan(&breadcrumbs)
  return breadcrumbs, result.Error
}

func GetProductCategoryListByParentAlternateKey(db *gorm.DB, AlternateKey string, ProductId uint) ( []ProductCategoryList, error) {
	var results []ProductCategoryList
	err := db.Table("categories AS parent").
      Select("child.name AS parameter_name, child.id AS category_id, parameters.id, parameters.markdown").
      Joins("INNER JOIN categories AS child ON child.parent_id = parent.id").
      Joins("LEFT JOIN product_parameters AS parameters ON child.id = parameters.category_id").
      Joins("LEFT JOIN products ON parameters.product_id = products.id AND products.id = ?", ProductId).
      Where("parent.alternate_key = ?", AlternateKey).
      Scan(&results).Error
	return results, err
}

*/
