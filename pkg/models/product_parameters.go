package models

/* ProductParameter Details could have been part of the products table, but for simpilicty,
   it has been placed as a single table with a 1:1 relationship to ProductParameter.  If we
   use this platform to sell a different product, then we can just edit this table */

import (
  "time"
//  "gorm.io/gorm"
//  "gorm.io/gorm/clause"
)

type ProductParameter struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    Markdown   string    `gorm:"default:null" json:"markdown"`
    CategoryID uint      `json:"category_id"`
    ProductID  uint      `json:"product_id"`

    // Association to Category
    Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`

    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}


/*
type ProductParameterItem struct {
    ParameterName string `gorm:"column:parameter_name" json:"parameter_name"`
    CategoryID    uint   `gorm:"column:category_id" json:"category_id"`
    ParameterID   uint   `gorm:"column:parameter_id" json:"parameter_id"`
    Markdown      string `gorm:"column:markdown" json:"markdown"`
}

type ParameterResponse []ProductParameterItem
*/


/*
// Start of CRUD interface
func GetProductParameters(db *gorm.DB, item *[]ProductParameter) (err error) {
	err = db.Find(item).Error
	if err != nil {
		return err
	}
	return nil
}

// get account by id
func GetProductParameter(db *gorm.DB, item *ProductParameter, id uint) (err error) {
	err = db.Where("id = ?", id).First(item).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProductParameter(db *gorm.DB, item *ProductParameter) (err error) {
	err = db.Save(item).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProductParameters(db *gorm.DB, items []ProductParameter) (err error) {
	err = db.Save(items).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateProductParameter(db *gorm.DB, item *ProductParameter) (*ProductParameter, error) {
	err := db.Create(&item).Error
	if err != nil {
		return &ProductParameter{}, err
	}
	return item, nil
}

func DeleteProductParameter(db *gorm.DB, item *ProductParameter) (err error) {
	err = db.Delete(item).Error
	if err != nil {
		return err
	}
	return nil
}

// end of CRUD iterface

// Start of utility functions


func GetProductParametersByProductId(db *gorm.DB, item *[]ProductParameter, ProductId uint) (err error) {
	err = db.Where("category_id = ?", CategoryId).Error
	if err != nil {
		return err
	}
	return nil
}

// Retrieve user list with eager loading credit cards
func GetProductParameterByCategoryAlternate_Key(db *gorm.DB, alternate_key string) ([]ProductParameter, error) {
    var suppliers []ProductParameter
    err := db.Model(&ProductParameter{}).Preload("Categories").Find(&suppliers).Where("alternate_key = ?", alternate_key).Error
    return suppliers, err
}
*/
