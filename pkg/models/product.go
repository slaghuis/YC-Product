package models

/* ************************************************************************************
 * Usage Flow
 * 1.	Create Product → Add variants and options.
 * 2.	Generate SKUs → Use GenerateAllSKUs to create all combinations.
* 3.	Query SKUs → Fetch SKUs with their options for display or inventory management.
 * ************************************************************************************ */
import (
	"time"
	"strings"
	"gorm.io/gorm"
)

type ProductState uint

const (
	ProductStateDraft 			ProductState = 100 // Created, but content is dodge
	ProductStateSuspended 	ProductState = 200 // No longer for sale.  Hide from prospective customers and remove from baskets. (Tag in wish lists)
	ProductStateActive 		  ProductState = 300 // Ready for sale.  Published to prospective customers
	ProductStatePromoted 		ProductState = 400 // On Promotion.  Published to prospective customers
)

/* Represents the base product (e.g., "T-Shirt"). */
type Product struct {
    ID          uint           `gorm:"primaryKey"`
    Name        string         `gorm:"size:255;not null"`
    Description string         `gorm:"type:text"`
    State       ProductState   `gorm:"default:100"`
		SupplierID  uint		 			 `gorm:"default:null"`
		Supplier    Supplier       `gorm:"foreignKey:SupplierID"` // <-- add this line
		Images 			[]Image        `gorm:"foreignKey:ProductID"`
    Variants    []Variant      `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
    SKUs        []SKU          `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
		Reviews  		[]Review       `gorm:"foreignKey:ProductID"`
		Likes  			[]Like     		 `gorm:"foreignKey:ProductID"`
		Categories      []Category         `gorm:"many2many:product_categories;"`
		ProductParameter 	[]ProductParameter  `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

/* Represents a dimension of variation (e.g., "Color", "Size"). */
type Variant struct {
    ID        uint       `gorm:"primaryKey"`
    ProductID uint       `gorm:"not null;index"`
    Name      string     `gorm:"size:255;not null"` // e.g., "Color"
    Options   []Option   `gorm:"foreignKey:VariantID;constraint:OnDelete:CASCADE;"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

/* Represents a possible value for a variant (e.g., "Red", "Blue", "Small"). */
type Option struct {
    ID        uint       `gorm:"primaryKey"`
    VariantID uint       `gorm:"not null;index"`
    Value     string     `gorm:"size:255;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time

		Variant   Variant    `gorm:"foreignKey:VariantID;constraint:OnDelete:CASCADE;"`
}

/*Represents a purchasable unit (product + chosen options). */
type SKU struct {
    ID        uint       `gorm:"primaryKey"`
    ProductID uint       `gorm:"not null;index"`
    Code      string     `gorm:"uniqueIndex;size:100"` // e.g., "TSHIRT-RED-SMALL"
    Price     float64        `gorm:"not null"`
    Stock     int        `gorm:"not null"`
		Options   []SKUOption `gorm:"foreignKey:SKUId;constraint:OnDelete:CASCADE;"`
		Product   Product     `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

/*Join table linking SKUs to chosen options. */
type SKUOption struct {
    ID      uint `gorm:"primaryKey"`
    SKUId   uint `gorm:"not null;index"`
    OptionID uint `gorm:"not null;index"`

		// Explicit parent relationship
    SKU    SKU    `gorm:"constraint:OnDelete:CASCADE;"`
    Option Option `gorm:"constraint:OnDelete:CASCADE;"`
}

func GetProductWithVariants(db *gorm.DB, productID uint) (*Product, error) {
    var product Product
    err := db.Preload("Variants.Options").
        Preload("SKUs.Options").
        First(&product, productID).Error
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func GetSKU(db *gorm.DB, skuID uint) (*SKU, error) {
    var sku SKU
    err := db.Preload("Options.Option").
        First(&sku, skuID).Error
    if err != nil {
        return nil, err
    }
    return &sku, nil
}


/* Build a SKU code from product name and chosen options. */
func GenerateSKUCode(product *Product, options []Option) string {
    parts := []string{strings.ToUpper(strings.ReplaceAll(product.Name, " ", ""))}
    for _, opt := range options {
        parts = append(parts, strings.ToUpper(opt.Value))
    }
    return strings.Join(parts, "-")
}

/* Generate SKU from Options */
func CreateSKU(db *gorm.DB, product Product, options []Option, price float64, stock int) (*SKU, error) {
    sku := SKU{
        ProductID: product.ID,
        Code:      GenerateSKUCode(&product, options),
        Price:     price,
        Stock:     stock,
    }

    if err := db.Create(&sku).Error; err != nil {
        return nil, err
    }

    for _, opt := range options {
        skuOpt := SKUOption{SKUId: sku.ID, OptionID: opt.ID}
        if err := db.Create(&skuOpt).Error; err != nil {
            return nil, err
        }
    }

    return &sku, nil
}


/* This function computes the Cartesian product of options across variants. */
func GenerateAllSKUs(db *gorm.DB, product Product, price float64, stock int) ([]SKU, error) {
    var skus []SKU

    // Load variants and options
    db.Preload("Variants.Options").First(&product)

    // Recursive helper to build combinations
    var build func(idx int, current []Option)
    build = func(idx int, current []Option) {
        if idx == len(product.Variants) {
            sku, err := CreateSKU(db, product, current, price, stock)
            if err == nil {
                skus = append(skus, *sku)
            }
            return
        }
        for _, opt := range product.Variants[idx].Options {
            build(idx+1, append(current, opt))
        }
    }

    build(0, []Option{})
    return skus, nil
}
