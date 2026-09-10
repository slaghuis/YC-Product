package dto

type ReviewListItem struct {
  Score       uint            `json:"score"`
  Comment     string          `json:"comment"`
  Moderated   bool            `json:"moderated"`
  FullName    string          `json:"full_name"`
  AccountID   string          `json:"account_id"`
  ProductName string          `json:"product_name"`
  ProductID   uint            `json:"product_id"`
}
