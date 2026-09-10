package models

import (
	"math"
	"strconv"
	"time"

	"gorm.io/gorm"
)


// Scope for filtering records thaat were created in the last day
func LastDay(db *gorm.DB) *gorm.DB {
	return db.Where("createdat >= ?", time.Now().AddDate(0,0,-1))
}

// Scope for limiting to the top 10
func Top10(db *gorm.DB) *gorm.DB {
	return db.Limit(10)
}

// Scope for pagination
func Paginate(value interface{}, pagination *Pagination) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        var totalRows int64
        // Count total rows for the given model
        db.Model(value).Count(&totalRows)

        pagination.TotalRows = totalRows
        pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

        // Apply pagination and sorting
        return db.Model(value).
            Offset(pagination.GetOffset()).
            Limit(pagination.GetLimit()).
            Order(pagination.GetSort())
    }
}

// Filter on categories.id
func FilterByCategoryID(categoryID string) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        // If blank, ignore
        if categoryID == "" {
            return db
        }

        // Try to parse as integer
        id, err := strconv.Atoi(categoryID)
        if err != nil {
            // Not a number → ignore filter
            return db
        }

        // Apply filter
        return db.Where("categories.id = ?", id)
    }
}
