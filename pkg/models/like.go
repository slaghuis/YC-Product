package models

import (
  "time"
)

type Like struct {
  ID          uint            `gorm:"primaryKey";json:"id"`
	Score       uint            `gorm:"default:1";json:"score"`
  AccountID   string          `json:"account_id"`
  ProductID   uint            `json:"product_id"`

  CreatedAt   time.Time       `json:"created_at"`
  UpdatedAt   time.Time       `json:"updated_at"`
}
