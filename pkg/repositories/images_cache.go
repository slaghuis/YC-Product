package repositories

import (
    "gorm.io/gorm"
    "time"
    "github.com/slaghuis/YC-Product/pkg/models"
)

type ImageCacheRepository struct {
    db *gorm.DB
}

func NewImageCacheRepository(db *gorm.DB) *ImageCacheRepository {
    return &ImageCacheRepository{db: db}
}

func (r *ImageCacheRepository) FindRecentImage(referenceId, referenceType string) (*models.ImageCache, error) {
    var cache models.ImageCache
    err := r.db.Where("reference_id = ? AND reference_type = ? AND expires_at < ?",
        referenceId, referenceType, time.Now()).
        First(&cache).Error
    if err != nil {
        return nil, err
    }
    return &cache, nil
}

func (r *ImageCacheRepository) SaveImage(cache *models.ImageCache) error {
    return r.db.Save(cache).Error
}
