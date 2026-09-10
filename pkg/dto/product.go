package dto

import "github.com/slaghuis/YC-Product/pkg/models"

// Just a simple structure to read from the database
type ProductListItem struct {
		ID          			uint
    Name        			string
		Promoted					bool
		Like							bool
}

// Exteded structure with content from the API's
type ProductListItemExtended struct {
		ID          			uint           			`json:"id"`
    Name        			string         			`json:"name"`
		Promoted					bool								`json:"promoted"`
		Like							bool								`json:"like"`
		ImagePath         string              `json:"image_path"`       // Read from imaging cache
		CheapestSKU     	SKUListItem         `json:"cheapest_sku"`
		Price							float64             `json:"price"`
		Currency          string              `json:"currency"`
		Variants				  []VariantGroup      `json:"variants"`
}

type ProductMerchantListItem struct {
		ID          			uint           			`json:"id"`
    Name        			string         			`json:"name"`
    Description 			string         			`json:"description"`
    State       			models.ProductState `json:"state"`
		ImagePath     		string 							`json:"image_path"`
		MinPrice     			float64 						`json:"min_price"`
		MaxPrice     			float64 						`json:"max_price"`
		Likes             int                 `json:"like_count"`
		Reviews           int                 `json:"review_count"`
		Sales							int                 `json:"sales"`
		Earnings					float64							`json:"earnings"`
}

type ProductParameterItem struct {
  ParameterName  string   `json:"parameter_name"`
  CategoryId     uint     `json:"category_id"`
  ParameterId    uint     `json:"parameter_id"`
  Markdown       string   `json:"markdown"`
}

type ProductBasicInfo struct {
	ID             uint                    `json:"id"`
	Name           string                  `json:"name"`
	Description    string                  `json:"description"`
	State          uint  	                 `json:"state"`
	SupplierID     uint                    `json:"supplier_id"`
	SupplierName   string                  `json:"supplier_name"`
	LikeCount      int                     `json:"like_count"`
	ReviewCount    int                     `json:"review_count"`
}

/*
type ProductMerchantDetail struct {
	ID             uint                    `json:"id"`
	Name           string                  `json:"name"`
	Description    string                  `json:"description"`
	State          uint  	                 `json:"state"`
	SupplierID     uint                    `json:"supplier_id"`
	SupplierName   string                  `json:"supplier_name"`
	LikeCount      int                     `json:"like_count"`
	ReviewCount    int                     `json:"review_count"`
	Sales          int                     `json:"sales"`
	Earnings       float64                 `json:"earnings"`
	Parameters     []ProductParameterItem	 `json:"parameters"`
  CategoryGroups []CategoryGroup         `json:"category_groups"`
  Variants       []VariantGroup          `json:"variants"`
}
*/

//THESE ARE NEW

type ProductResponse struct {
    ProductID    uint                `json:"product_id"`
    ProductName  string              `json:"product_name"`
    Description  string              `json:"description"`
    SupplierID   uint                `json:"supplier_id"`
    SupplierName string             `json:"supplier_name"`
    Promoted     bool                `json:"promoted"`
    Price        float64             `json:"price"`
    Variants     []VariantResponse   `json:"variants"`
    Parameters   []ParameterResponse `json:"parameters"`
    Categories   []CategoryResponse  `json:"categories"`
}

type VariantResponse struct {
    Name    string           `json:"name"`
    Options []OptionResponse `json:"options"`
}

type OptionResponse struct {
    OptionID   uint   `json:"option_id"`
    OptionName string `json:"option_name"`
    Selected   bool   `json:"selected"`
}

type ParameterResponse struct {
    ID        uint   `json:"id"`
    Name      string `json:"name"`
    Markdown  string `json:"markdown"`
}

type CategoryResponse struct {
    ID           uint   `json:"id"`
    Name         string `json:"name"`
    AlternateKey string `json:"alternate_key"`
    IconPath     string `json:"icon_path"`
}
