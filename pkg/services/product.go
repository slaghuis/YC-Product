package services

import (
    "fmt"
    "context"
    "encoding/json"
    "errors"
    "math"
    "sort"

//    "github.com/slaghuis/YC-Product/pkg/services"
    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
    "github.com/slaghuis/YC-Product/pkg/repositories"
)

type productService struct {
    productRepo       *repositories.ProductRepository
    cacheRepo         *repositories.ImageCacheRepository
    skuService        SkuServiceInterface
    productSearchService ProductSearchService
    preferenceClient  PreferenceClient
    imagesClient      ImagesClient
}

func NewProductService(
  productRepo *repositories.ProductRepository,
  cacheRepo   *repositories.ImageCacheRepository,
  searchSvc  ProductSearchService,
  skuService SkuServiceInterface,
  preference PreferenceClient,
  images ImagesClient,
  ) ProductService {
    return &productService{
      productRepo: productRepo,
      cacheRepo: cacheRepo,
      skuService: skuService,
      productSearchService: searchSvc,
      preferenceClient: preference,
      imagesClient: images,
    }
}

func (s *productService) ReadProductCategories(ctx context.Context, id string) (dto.CategoriesResponse, error) {
  // Now read the product Categories in two steps.  Find all the category types
  // and then list the categories that we have to deal with
  baseKeys, err := s.productRepo.ReadBaseCategories()
  if err != nil {return dto.CategoriesResponse{}, err }

  // Prepare response
  category_response := dto.CategoriesResponse{
      Groups: make([]dto.CategoryGroup, 0, len(baseKeys)),
  }

  // Loop through base keys and run second query for each
  for _, bk := range baseKeys {
      productCategories, err := s.productRepo.ListProductCategoryItem(bk.AlternateKey, id)
      if  err != nil {
          return category_response, err
      }

      // Add group with categories
      category_response.Groups = append(category_response.Groups, dto.CategoryGroup{
          AlternateKey: bk.AlternateKey,
          GroupName:    bk.Name,
          Categories:   productCategories,
      })
  }

  return category_response, nil
}

func (s *productService) UpdateProductCategories(ctx context.Context, productID uint, categoryIDs []uint) (error) {
  return s.productRepo.UpdateProductCategories( productID, categoryIDs)
}

func (s *productService) ReadProductParameters(ctx context.Context, productID string) ([]dto.ProductParameterItem, error) {
  return s.productRepo.ReadProductParameters(productID)
}

func (s *productService) SaveProductParameters(ctx context.Context, productParameters []models.ProductParameter) (error) {
  return s.productRepo.SaveProductParameters(productParameters)
}

func (s *productService) CreateProduct(ctx context.Context, name, description string, supplierID uint) (*models.Product, error) {
  item := &models.Product {
    Name        : name,
    Description : description,
    SupplierID  : supplierID,
    State			  : models.ProductStateDraft,
  }
  err := s.productRepo.SaveProduct(item)
  return item, err
}

func (s *productService) ListProducts(ctx context.Context, page, limit, sort, category, token string) (*models.Pagination, error) {
  pagination := &models.Pagination{
    Page: StrToInt(page,1),
    Limit: StrToInt(limit,20),
//    Sort: sort,
  }

  // Call the preference service to get the search criteria
  preferences, err := s.preferenceClient.GetUserPreferences(token)
  if err != nil {
    fmt.Printf("Cannot get the user preferences. %s\n", err)
    return pagination, err
  }
  order, err := s.preferenceClient.GetUserSortOrder(token)
  if err != nil {
    pagination.Sort = "relevance"
  } else {
    pagination.Sort = order.SortOrder
  }
  fmt.Printf("Sort order will be %s\n", pagination.Sort)

  //include the categroy filter from the URL
  if len(category)>0 {
    categoryFilter := dto.Filter{
      FilterType: "category",
      Values: []string{category},
    }
    preferences.Filters = append(preferences.Filters, categoryFilter)
  }

  // LogJson(preferences)

  // Call the search service to get the product IDs []int
  productIDs, err := s.productSearchService.SearchProducts(ctx, *preferences, pagination.Sort)
  if err != nil {
    fmt.Printf("Search Service filed to retreive products. %s\n", err)
    return pagination, err
  }

  pagination.TotalRows = int64(len(productIDs))

  productPage := Paginate(productIDs, pagination.Page, pagination.Limit)
  if len(productPage) == 0 {
    fmt.Printf("No records found for this page.\n")
    return pagination, errors.New("No records for this page")
  }

  var response []dto.ProductListItemExtended
  for _, p := range productPage {
    // For each product ...
    product , err := s.productRepo.GetProductAndCategories(p) //returns dto.ProductListItem, err
    if err != nil {
      continue
    }

    // Read the variants an options available
    variants, err := s.GetVariantList(ctx, IntToStr(p))
    if err != nil {
      continue
    }

    // Read the image path from the image cache
    imagepath, _ := s.GetImage(IntToStr(p), "product", token)

    // Read the pricing from the pricing cache
    // -- List all SKU's for this product
    skus, err := s.skuService.GetProductSKUList(ctx, IntToStr(p), token)
    if err != nil {
      continue
    }
    var cheapestPrice float64 = math.MaxFloat64
    var cheapest dto.SKUListItem

    for _, sku := range skus {  // typically 2 to 3 SKU's  This should be quick
      // -- Distill the minimum FinalPrice from here
      if (sku.Price > 0.0) && (sku.Price < cheapestPrice) {
        cheapestPrice = sku.Price
        cheapest = sku
      }
    }

    // Safeguard to excuse products that do not have pricing set.
    if cheapest.ID == 0 {
      continue
    }

    response = append(response, dto.ProductListItemExtended{
        ID:             product.ID,
        Name:           product.Name,
        Promoted:       product.Promoted,
        Like:           product.Like,
        ImagePath:      imagepath,
        CheapestSKU:    cheapest,
        Price:          cheapestPrice,
        Variants:			  variants,
    })

  }
  pagination.TotalPages = int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.Limit)))
  pagination.Rows = response
  pagination.Context = preferences//fmt.Sprintf("category=%s",category)   //Idea here is to say "category=7" or "vegetarian"

  return pagination, err
}


func (s *productService) CountProducts(ctx context.Context, category, token string) (*dto.ProductCount, error) {
  // Call the preference service to get the search criteria
  preferences, err := s.preferenceClient.GetUserPreferences(token)
  if err != nil {
    return nil, err
  }

  //include the categroy filter from the URL
  if len(category)>0 {
    categoryFilter := dto.Filter{
      FilterType: "category",
      Values: []string{category},
    }
    preferences.Filters = append(preferences.Filters, categoryFilter)
  }

  // Call the search service to get the product IDs
  productIDs, err := s.productSearchService.SearchProducts(ctx, *preferences, "")
  if err != nil {
    return nil, err
  }

  result := dto.ProductCount {
    Count: len(productIDs),
    Preferences: *preferences,
  }

  return &result, nil
}


func (s *productService) ListMerchantProducts(ctx context.Context, page, limit, sort string) (*models.Pagination, error) {
  pagination := &models.Pagination{
    Page: StrToInt(page,1),
    Limit: StrToInt(limit,20),
    Sort: sort,
  }

  productList, err := s.productRepo.GetMerchantProdictList(pagination)

  pagination.Rows = productList
  pagination.Context = "merchant"   //Idea here is to say "category=7" or "vegetarian"

  return pagination, err
}

func (s *productService) GetProductBasicInfo(ctx context.Context, id string) (*dto.ProductBasicInfo, error) {
  return s.productRepo.GetProductBasicInfo(id)
}

func (s *productService) UpdateProduct(ctx context.Context, id, name, description string, state models.ProductState) (*models.Product, error) {
  productID, err := StrToUint(id)
  if err != nil {
    return nil, err
  }

  item, err := s.productRepo.GetProduct(productID)

  // Overwrite all except the ID
  item.Name         = name
  item.State      	= state
  item.Description  = description
// Mmmm,  I cannot change the supplier yet?

 err = s.productRepo.SaveProduct(item)
 return item, err
}

func (s *productService) SetState(ctx context.Context, id, state string) (error) {
  productID, err := StrToUint(id)
  if err != nil {
    return nil
  }

  newState := StrToInt(state, 100)

  return s.productRepo.SetState(productID, newState)
}


func (s *productService) GetShopProduct(ctx context.Context, productID string) (*dto.ProductShopListItem, error) {

  row, err :=  s.productRepo.GetShopProduct(productID)
  if err != nil { return nil, err }
  // Convert scan rows -> response items
   product := dto.ProductShopListItem{
     ID          : row.ID,
     Name        : row.Name,
     Description : row.Description,
     Promoted    : row.Promoted,
     ImagePath:    row.ImagePath,
     MinPrice:     row.MinPrice,
     MaxPrice:     row.MaxPrice,
     LikesCount:   row.LikesCount,
     ReviewsCount: row.ReviewsCount,
     SKUCount:     row.SKUCount,
   }

   //Unmarshal the raw JSON into the typed slice
   if row.SKUsJSON != "" && row.SKUsJSON != "[]" {
     if err := json.Unmarshal([]byte(row.SKUsJSON), &product.SKUs); err != nil {
       return nil, err
     }
   } else {
     product.SKUs = []dto.SKUReturnItem{}
   }

   return &product, nil
}


func (s *productService)GetShopProductDetail(ctx context.Context, id string) (*dto.ProductResponse, error) {
  productID, err := StrToUint(id)
  if err != nil { return nil, err }

  product, err := s.productRepo.GetShopProductDetail(productID)

  // sort parameters by category sort order
  sort.Slice(product.ProductParameter, func(i, j int) bool {
      return product.ProductParameter[i].Category.SortOrder <
             product.ProductParameter[j].Category.SortOrder
  })

  response := dto.ProductResponse{
    ProductID:   product.ID,
    ProductName: product.Name,
    Description: product.Description,
    SupplierID:   product.Supplier.ID,
    SupplierName: product.Supplier.Name,
    Promoted   : (product.State == models.ProductStatePromoted),
    Price:       24.99,
    Variants:    []dto.VariantResponse{},
    Parameters:   []dto.ParameterResponse{},
    Categories:   []dto.CategoryResponse{},
  }

  // map variants
  for _, v := range product.Variants {
    variantResp := dto.VariantResponse{
        Name:    v.Name,
        Options: []dto.OptionResponse{},
    }
    for i, o := range v.Options {
        variantResp.Options = append(variantResp.Options, dto.OptionResponse{
            OptionID:   o.ID,
            OptionName: o.Value,
            Selected: (i==0),
        })
    }
    response.Variants = append(response.Variants, variantResp)
  }

  // map parameters
  for _, p := range product.ProductParameter {
      paramResp := dto.ParameterResponse{
          ID:          p.ID,
          Name:        p.Category.Name,
          Markdown:    p.Markdown,
      }
      response.Parameters = append(response.Parameters, paramResp)
  }

  // map categories
  for _, c := range product.Categories {
      response.Categories = append(response.Categories, dto.CategoryResponse{
          ID:           c.ID,
          Name:         c.Name,
          AlternateKey: c.AlternateKey,
          IconPath:     c.IconPath,
      })
  }

  return &response, nil
}

func (s *productService)GetProductFromSKU(ctx context.Context, skuCode, token string) (*dto.ProductVariantDetail,error) {

  //If this does not return the variants and other stuff then check out the pricing cache!
  sku, err := s.skuService.GetSKU(ctx, skuCode, token)
  if err != nil { return nil, err }

  variantMap := make(map[uint]*dto.VariantResponse)

  for _, skuOpt := range sku.Options {
      variant := skuOpt.Option.Variant
      option := skuOpt.Option

      // If we haven’t seen this variant yet, create it
      if _, exists := variantMap[variant.ID]; !exists {
          variantMap[variant.ID] = &dto.VariantResponse{
              Name:    variant.Name,
              Options: []dto.OptionResponse{},
          }
      }

      // Append the option to the variant
      variantMap[variant.ID].Options = append(
          variantMap[variant.ID].Options,
          dto.OptionResponse{
              OptionID:   option.ID,
              OptionName: option.Value,
              Selected:   true, // since this SKU has chosen this option
          },
      )
  }

  // Flatten map into slice
  variants := make([]dto.VariantResponse, 0, len(variantMap))
  for _, v := range variantMap {
      variants = append(variants, *v)
  }

  item := dto.ProductVariantDetail{
    ID : sku.Product.ID,
    Name: sku.Product.Name,
    Variants: variants,
  }

  return &item, nil
}
