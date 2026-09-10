package models

import (
  "time"
)

type Review struct {
  ID          uint            `gorm:"primaryKey";json:"id"`
	Score       uint            `gorm:"default:5";json:"score"`
  Comment     string          `json:"comment"`
	Moderated   bool            `gorm:"default:false";json:"moderated"`
  FullName    string          `gorm:"default:Anonymous";json:"full_name"`
  AccountID   string          `json:"account_id"`
  ProductID   uint            `json:"product_id"`

  CreatedAt   time.Time       `json:"created_at"`
  UpdatedAt   time.Time       `json:"updated_at"`
}
