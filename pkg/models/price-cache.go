package models

import (
  "gorm.io/gorm"
  "time"
  "github.com/google/uuid"
)

type PriceCache struct {
    CacheID        string    `gorm:"type:uuid;primaryKey"`
    SkuCode        string    `gorm:"index"`
    RegionCode     string    `gorm:"index"`
    CustomerSegment string   `gorm:"index"`
    CurrencyCode   string    `gorm:"index"`
    OriginalPrice  float64
    FinalPrice     float64
    AppliedRules   string    // JSON or comma-separated list
    CachedAt       time.Time
    ExpiresAt      time.Time
}

func (b *PriceCache) BeforeCreate(tx *gorm.DB) (err error) {
    if b.CacheID == "" {
        b.CacheID = uuid.NewString()
    }
    return
}
