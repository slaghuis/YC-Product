package dto

import (
  "encoding/json"
  "fmt"
)

type OptionValueItem struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}

type OptionValueItemList []OptionValueItem

func (s *OptionValueItemList) Scan(src interface{})  error {
  if src == nil {
    *s = []OptionValueItem{}
    return nil
  }
  switch v := src.(type) {
  case []byte:
    return json.Unmarshal(v, s)
  case string:
    return json.Unmarshal([]byte(v), s)
  default:
    return fmt.Errorf("Unsupported type for OptionValueItemList: :T", src)
  }
}

type SKUReturnItem struct {
    ID      uint                `json:"id"`
    Code    string              `json:"code"`
    Price   float64             `json:"price"`
    Stock   int                 `json:"stock"`
    Options OptionValueItemList `json:"options"`
}

type SKUList []SKUReturnItem

func (s *SKUList) Scan(src interface{})  error {
  if src == nil {
    *s = []SKUReturnItem{}
    return nil
  }
  switch v := src.(type) {
  case []byte:
    return json.Unmarshal(v, s)
  case string:
    return json.Unmarshal([]byte(v), s)
  default:
    return fmt.Errorf("Unsupported type for SKUList: :T", src)
  }
}

type ProductShopListItem struct {
    ID           uint             `json:"id"`
    Name         string           `json:"name"`
    Description  string           `json:"description"`
    Promoted     bool             `json:"promoted"`
    ImagePath    string           `json:"image_path"`
    MinPrice     float64          `json:"min_price"`
    MaxPrice     float64          `json:"max_price"`
//    LikesCount   int              `json:"likes_count"`
//    ReviewsCount int              `json:"reviews_count"`
    SKUCount     int              `json:"sku_count"`
    SKUs         SKUList          `json:"skus"`
}

type productScanRow struct {
    ID           uint             `gorm:"column:id"`
    Name         string           `gorm:"column:name"`
    Description  string           `gorm:"column:description"`
    Promoted     bool             `gorm:"column:promoted"`
    ImagePath    string           `gorm:"column:image_path"`
    MinPrice     float64          `gorm:"column:min_price"`
    MaxPrice     float64          `gorm:"column:max_price"`
//    LikesCount   int              `gorm:"column:likes_count"`
//    ReviewsCount int              `gorm:"column:reviews_count"`
    SKUCount     int              `gorm:"column:sku_count"`
    SKUsJSON     string           `gorm:"column:skus"`
}
