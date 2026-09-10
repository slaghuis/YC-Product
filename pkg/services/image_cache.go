package services

import (
  "time"
  "github.com/google/uuid"
  "github.com/slaghuis/YC-Product/pkg/models"
)

func  (s *productService) GetImage(referenceId, referenceType, token string) (string, error) {
  var response string

  // Check cache first
  cached, _ := s.cacheRepo.FindRecentImage(referenceId, referenceType)
  if cached != nil {
    response = cached.Thumbnail
  } else {
      resp, err := s.imagesClient.GetPrimaryImage(referenceId, referenceType, token)
      if err != nil {
          // Return a default "No Image" type image
          // Load this from configuratio
          response = ""
      }

      // Save to cache
      s.cacheRepo.SaveImage(&models.ImageCache{
          CacheID:        uuid.NewString(),
          Category:       resp.Category,
          Original:       resp.Original,
          Main:           resp.Main,
          Thumbnail:      resp.Thumbnail,
          ReferenceID:    resp.ReferenceID,
          ReferenceType:  resp.ReferenceType,
          SortOrder:      resp.SortOrder,
          IsPrimary:      resp.IsPrimary,
          CachedAt:       time.Now(),
          ExpiresAt:      time.Now().Add(15 * time.Minute),
      })

      response = resp.Thumbnail
  }
  return response, nil
}
