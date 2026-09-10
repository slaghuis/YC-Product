package repositories

import (
    "github.com/k-samuel/faceted/search"
//    "github.com/slaghuis/YC-Product/pkg/models"
)

// Faceted-backed repository
type SearchRepository struct {
    db *search.Db
}

type SearchRecord map[string]interface{}

func NewSearchRepository() *SearchRepository {
  return &SearchRepository{
    db: search.NewContainer().NewDb(),
  }
}

func (r *SearchRepository) Insert(recordId int, item SearchRecord) {

  r.db.GetStorage().AddRecord(recordId, item)
//  recordId := int(item["id"].(int))
//  delete(item, "id")
//  r.storage.AddRecord(recordId, item)
}


func (r *SearchRepository) Optimize() {
  r.db.GetStorage().Optimize()
}



func (r *SearchRepository) Query(query *search.SearchQuery) ([]int, error) {
    return r.db.Query(query)
}

/*

func (r *SearchRepository) IndexProducts(ctx context.Context) error {
    // Load products from GORM
    var products []models.Product
    if err := r.productRepo.db.Find(&products).Error; err != nil {
        return err
    }

    // Index each product into faceted
    for _, p := range products {
        record := search.NewRecord(p.ID).
            SetField("name", p.Name).
            SetField("description", p.Description).
            SetField("state", p.State).
            SetField("brand", p.Brand).
            SetField("category", p.Category).
            SetField("price", p.Price).
            SetField("color", p.Color).
            SetField("size", p.Size)

        r.searchDb.Insert(record)
    }

    return nil
}

func (r *ProductSearchRepository) QueryProducts(query *search.SearchQuery) ([]dto.ProductListItem, error) {
    records := r.searchDb.Query(query)
    var productList []dto.ProductListItem
    for _, rec := range records {
        productList = append(productList, dto.ProductListItem{
            ID:          rec.ID.(uint),
            Name:        rec.Fields["name"].(string),
            Description: rec.Fields["description"].(string),
            State:       rec.Fields["state"].(int),
        })
    }
    return productList, nil
}
*/
