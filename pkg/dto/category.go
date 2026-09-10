package dto

type CategoryMerchantListItem struct {
  AlternateKey  string   `json:"alternate_key"`
  ParentId      uint     `json:"parent_id"`
  Id            uint     `json:"id"`
  Name          string   `json:"name"`
  IconPath      string   `json:"icon_path"`
  Description   string   `json:"description"`
  LikesCount    uint     `json:"likes_count"`
  ReviewCount   uint     `json:"review_count"`
  ProductCount  int      `json:"product_count"`
  Sales					int      `json:"sales"`
  Earnings			float64	 `json:"earnings"`
}

type CategoryListItem struct {
  AlternateKey  string   `json:"alternate_key"`
  ParentId      uint     `json:"parent_id"`
  Name          string   `json:"name"`
  IconPath      string   `json:"icon_path"`
  SubType       uint     `json:"sub_type"`
  Description   string   `json:"description"`
  SortOrder     uint     `json:"sort_order"`
  ProductCount  int      `json:"product_count"`
}
