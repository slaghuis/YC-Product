package models

import "time"

type Supplier struct{
  ID            uint           `gorm:"primaryKey";json:"id"`
  Name          string         `gorm:"not null;default:null";json:"name"`
  RepName       string         `json:"rep_name"`
  LogoPath      string         `json:"logo_path"`
  Phone         string         `json:"phone"`
  PhoneVerified bool           `json:"phone_verified"`
  Email         string         `json:"email"`
  SMS           bool           `gorm:"default:false";json:"sms"`
  WhatsApp      bool           `gorm:"default:false";json:"whatsapp"`
  EmailNotify   bool           `gorm:"default:true";json:"email_notify"`
  AppPush       bool           `gorm:"default:false";json:"app_push"`
  Product       []Product      `gorm:"foreignKey:SupplierID";json:"products"`

  CreatedAt     time.Time      `json:"created_at"`
  UpdatedAt     time.Time      `json:"updated_at"`
}



// Start of utility functions
/*
func GetSuppliersByCategoryId(db *gorm.DB, item *[]Supplier, CategoryId uint) (err error) {
	err = db.Where("category_id = ?", CategoryId).Error
	if err != nil {
		return err
	}
	return nil
}

// Retrieve user list with eager loading credit cards
func GetSupplierByCategoryAlternate_Key(db *gorm.DB, alternate_key string) ([]Supplier, error) {
    var suppliers []Supplier
    err := db.Model(&Supplier{}).Preload("Categories").Find(&suppliers).Where("alternate_key = ?", alternate_key).Error
    return suppliers, err
}
*/
