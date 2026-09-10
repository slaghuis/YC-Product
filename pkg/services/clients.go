package services

import (
  "time"
  "github.com/slaghuis/YC-Product/pkg/dto"
)

type IdentityClient interface {
    GetContext(token string) (*IdentityContext, error)
}

type PricingClient interface {
    Calculate(req PricingRequest, token string) (PricingResponse, error)
    CreatePrice(req SetRequest, token string) (PricingSKU, error)
    UpdatePrice(req PricingUpdate, token string) (PricingSKU, error)
}

type ImagesClient interface {
    GetPrimaryImage(reference_id, reference_type, token string) (ImagesResponse, error)
}

type PreferenceClient interface {
    GetUserPreferences(token string) (*dto.Preferences, error)
    GetUserSortOrder(token string) (*dto.SortOrder, error)
}

type IdentityContext struct {
    RegionCode      string  `json:"region_code"`
    CustomerSegment string  `json:"customer_segment"`
    Currency        string  `json:"currency_code"`  // As per ISO 4217
}

type PricingRequest struct {
    SkuCode         string    `json:"sku_code"`
    Quantity        int       `json:"quantity"`
    RegionCode      string    `json:"region_code"`
    CustomerSegment string    `json:"customer_segment"`
    TargetCurrency  string    `json:"target_currency"`
}


type PricingUpdate struct {
    Code         string  `json:"code"`
    Name         string  `json:"name"`
    BasePrice    float64 `json:"base_price"`
    CurrencyCode string  `json:"currency_code"`
}

// PricingResponse mirrors your pricing microservice contract
type PricingResponse struct {
    OriginalPrice float64   `json:"original_price"`
    FinalPrice    float64   `json:"final_price"`
    CurrencyCode  string    `json:"currency_code"`
    AppliedRules  []uint    `json:"applied_rules"`
    CalculatedAt  time.Time `json:"calculated_at"`
}

type SetRequest struct {
    Code         string  `json:"code" binding:"required"`
    Name         string  `json:"name"`
    BasePrice    float64 `json:"base_price" binding:"required"`
    CurrencyCode string  `json:"currency_code" binding:"required"`
}

type PricingSKU struct {
    ID           uint           `json:"id"`
    Code         string         `json:"code"`
    Name         string         `json:"name"`
    BasePrice    float64        `json:"base_price"`
    CurrencyCode string         `json:"currency_code"`
    CreatedAt    time.Time      `json:"created_at"`
    UpdatedAt    time.Time      `json:"updated_at"`
    DeletedAt    time.Time      `json:"deleted_at"`
}

type ImagesResponse struct {
  ID            string     `json:"id"`
  Category      string     `json:"category"`
  Original      string     `json:"original"`
  Main          string     `json:"main"`
  Thumbnail     string     `json:"thumbnail"`
  ReferenceID   *string    `json:"reference_id"`     // product_id, caregory_id ...
  ReferenceType *string    `json:"reference_type"`   // product, category, promotion, blog
  SortOrder     int        `json:"sort_order"`
  IsPrimary     bool       `json:"primary"`
  CreatedAt     time.Time  `json:"created_at"`
  UpdatedAt     time.Time  `json:"updated_at"`
}
