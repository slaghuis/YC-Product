package repositories

import (
    "gorm.io/gorm"
    "time"
    "github.com/slaghuis/YC-Product/pkg/models"
)

type PriceCacheRepository struct {
    db *gorm.DB
}

func NewPriceCacheRepository(db *gorm.DB) *PriceCacheRepository {
    return &PriceCacheRepository{db: db}
}

func (r *PriceCacheRepository) FindValidPrice(sku, region, segment, currency string) (*models.PriceCache, error) {
    var cache models.PriceCache
    err := r.db.Where("sku_code = ? AND region_code = ? AND customer_segment = ? AND currency_code = ? AND expires_at > ?",
        sku, region, segment, currency, time.Now()).
        First(&cache).Error
    if err != nil {
        return nil, err
    }
    return &cache, nil
}

func (r *PriceCacheRepository) SavePrice(cache *models.PriceCache) error {
    return r.db.Save(cache).Error
}
