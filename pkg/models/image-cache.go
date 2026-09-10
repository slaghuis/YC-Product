package models

import (
  "gorm.io/gorm"
  "time"
  "github.com/google/uuid"
)

type ImageCache struct {
    CacheID        string    `gorm:"type:uuid;primaryKey"`
    Category      string     `gorm:"type:varchar(50);index" json:"category"`
    Original      string     `json:"original"`
    Main          string     `json:"main"`
    Thumbnail     string     `json:"thumbnail"`
    ReferenceID   *string    `gorm:"index" json:"reference_id,omitempty"`     // product_id, caregory_id ...
    ReferenceType *string    `gorm:"index" json:"reference_type,omitempty"`   // product, category, promotion, blog
    SortOrder     int        `gorm:"default:0" json:"sort_order"`
    IsPrimary     bool       `gorm:"default:false" json:"primary"`
    CachedAt       time.Time
    ExpiresAt      time.Time
}

func (b *ImageCache) BeforeCreate(tx *gorm.DB) (err error) {
    if b.CacheID == "" {
        b.CacheID = uuid.NewString()
    }
    return
}
