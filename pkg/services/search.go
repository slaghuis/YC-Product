package services

import (
  "context"

  "github.com/k-samuel/faceted/search"

  "github.com/slaghuis/YC-Product/pkg/dto"
  "github.com/slaghuis/YC-Product/pkg/repositories"
)

type productSearchService struct {
    productRepo     *repositories.ProductRepository
    searchRepo      *repositories.SearchRepository
}

func NewProductSearchService(productRepo *repositories.ProductRepository, searchRepo *repositories.SearchRepository) ProductSearchService {
    return &productSearchService{
        productRepo: productRepo,
        searchRepo:  searchRepo,
    }
}

// Index products from GORM into faceted
func (s *productSearchService) IndexProducts(ctx context.Context) error {

    products, err := s.productRepo.GetProductIndexProjection()
    if err != nil {return err}

    for _, p := range products {
      // TODO  Add the price to this search.
      // skuList, err := s.SkuService.GetProductSKUList(ctx, p.ID, token)
      // Find the minimum price for the SKU's

      // Cast pq.StringArray as []interface{} through iteration
      categoryInterfaceSlice := make([]interface{}, len(p.Categories))
      for i, v := range p.Categories {
        categoryInterfaceSlice[i] = v
      }

      variantInterfaceSlice := make([]interface{}, len(p.OptionValues))
      for i, v := range p.OptionValues {
        variantInterfaceSlice[i] = v
      }

      record := repositories.SearchRecord {  //map[string]interface{}
        "promotion" : p.Promoted,
        "supplier"   : p.Supplier,
        "category"  : categoryInterfaceSlice, //[]interface{}{"Vegan","Easter"},
        "variants"  : variantInterfaceSlice,
        "price": 350,
        "orders": 30,
        "age": p.Age,
      }

      s.searchRepo.Insert(int(p.ProductID), record)

    }

//    s.searchRepo.Optimize()

    return nil
}

// Search products using preferences
func (s *productSearchService) SearchProducts(ctx context.Context, prefs dto.Preferences, order string) ([]int, error) {
    filters := mapPreferencesToFacetedFilters(&prefs)

    var fieldName string
	  var direction, sortType int
    switch order {
    case "priceLH" :
      fieldName = "price"
      direction = search.SortAsc
      sortType = search.SortTypeNumbers
    case "priceHL" :
      fieldName = "price"
      direction = search.SortDesc
      sortType = search.SortTypeNumbers
    case "popularity" :
      fieldName = "orders"
      direction = search.SortDesc
      sortType = search.SortTypeNumbers
    case "arrival" :
      fieldName = "age"
      direction = search.SortDesc
      sortType = search.SortTypeNumbers
    default :  //"relevance"
      fieldName = ""
    }

    query := search.NewSearchQuery().Filters(filters)
    if fieldName != "" {
      query = query.Sort(fieldName, direction, sortType)
    }

    return s.searchRepo.Query(query)  //Records is a []int
}

/*
func (s *ProductSearchService) SearchProducts(prefs dto.Preferences) ([]dto.ProductListItem, error) {
    filters := mapPreferencesToFacetedFilters(&prefs)
    query := search.NewSearchQuery().Filters(filters)

    records, err := s.searchRepo.Query(query)  //Records is a []int
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

func mapPreferencesToFacetedFilters(prefs *dto.Preferences) []search.FilterInterface {
    var filters []search.FilterInterface

    for _, f := range prefs.Filters {
        switch f.FilterType {
        case "promotion", "supplier", "category", "variant":
            if len(f.Values) == 0 { continue }

            vals := make([]interface{}, len(f.Values))
            for i, v := range f.Values {
                vals[i] = v
            }
            filters = append(filters, search.NewValueFilter(f.FilterType, vals))

        case "priceRange":
            if f.Range != nil {
                filters = append(filters,
                    search.NewRangeFilter("price",
                        search.NewRangeValue(f.Range.Min, f.Range.Max)))
            }
        }
    }

    return filters
}
