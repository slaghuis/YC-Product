package dto


type ProductCategory struct {
    ProductID  uint `gorm:"primaryKey"`
    CategoryID uint `gorm:"primaryKey"`
}

// First query result
type BaseCategory struct {
    AlternateKey string `gorm:"column:alternate_key" json:"alternate_key"`
    Name         string `gorm:"column:name" json:"group_name"`
}

// Second query result
type ProductCategoryItem struct {
    CategoryID   int    `gorm:"column:category_id" json:"category_id"`
    CategoryName string `gorm:"column:category_name" json:"category_name"`
    IconPath     string `gorm:"column:icon_path" json:"icon_path"`
    Selected     bool   `gorm:"column:selected" json:"selected"`
}

// Final JSON response
type CategoriesResponse struct {
    Groups []CategoryGroup `json:"groups"`
}

// Each group contains its key, name, and categories
type CategoryGroup struct {
    AlternateKey string            			`json:"alternate_key"`
    GroupName    string            			`json:"group_name"`
    Categories   []ProductCategoryItem 	`json:"categories"`
}
