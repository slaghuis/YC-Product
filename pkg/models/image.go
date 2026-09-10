package models

/* ProductParameter Details could have been part of the products table, but for simpilicty,
   it has been placed as a single table with a 1:1 relationship to ProductParameter.  If we
   use this platform to sell a different product, then we can just edit this table */

import "time"

type ImageType uint

const (
	ImageTypeProduct 				ImageType = 100 // Main, high resolution product image
	ImageTypeThumbnail 		  ImageType = 200 // Thumbnails used for product catalogue
	ImageTypeGallery 				ImageType = 300 // Additoinal high resolution images
)

type Image struct {
	ID                uint            `gorm:"primaryKey" json:"id"`
	Name    					string          `gorm:"default:null" json:"name"`
  AltText  					string					`json:"alt_text"`
	ImagePath      		string         	`json:"image_path"`
	Type							ImageType				`gorm:"default:100" json:"type"`
  ProductID   			uint						`json:"product_id"`

	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}
