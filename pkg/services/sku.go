package services

import (
    "context"
    "math"
    "time"

    "github.com/google/uuid"

    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
    "github.com/slaghuis/YC-Product/pkg/repositories"
    "github.com/slaghuis/YC-Product/pkg/utils"
)

type SkuService struct {
    skuRepo        *repositories.SkuRepository
    cacheRepo      *repositories.PriceCacheRepository
    productRepo    *repositories.ProductRepository
    identityClient IdentityClient
    pricingClient  PricingClient
}

func NewSkuService(
    skuRepo        *repositories.SkuRepository,
    cacheRepo      *repositories.PriceCacheRepository,
    productRepo    *repositories.ProductRepository,
    identity       IdentityClient,
    pricing        PricingClient) SkuServiceInterface {
    return &SkuService{
      skuRepo: skuRepo,
      cacheRepo: cacheRepo,
      productRepo: productRepo,
      identityClient: identity,
      pricingClient: pricing,
    }
}

func (s *SkuService) UpdateSKU(ctx context.Context, id string, price float64, stock int, token string) (*models.SKU, error) {

  item, err := s.skuRepo.FindByID(id)
  if err != nil {return nil, err}

  item.Price = price
  item.Stock = stock
  err = s.skuRepo.SaveSKU(item)
  if err != nil {return nil, err}

  // Upddate the pricing API with the new price
  req := PricingUpdate {
      Code         : item.Code,
      Name         : "",
      BasePrice    : price,
      CurrencyCode : "ZAR",
  }
  _, err = s.pricingClient.UpdatePrice(req, token)

  return item, err
}

func (s *SkuService) CreateSKU(ctx context.Context, id string, OptionIDs []uint, price float64, stock int, token string) (*models.SKU, error) {
  productID, err := StrToUint(id)
  if err != nil {
    return nil, err
  }

  product, err := s.productRepo.GetProduct(productID)
  if err != nil {
    return nil,err
  }

  options, err := s.productRepo.GetOptions(OptionIDs)
  if err != nil {
    return nil,err
  }

  sku := models.SKU{
      ProductID: product.ID,
      Code:      models.GenerateSKUCode(product, options),
      Price:     price,
      Stock:     stock,
  }

  err = s.skuRepo.SaveSKU(&sku)
  if err != nil {
    return nil,err
  }

  for _, opt := range options {
      skuOpt := models.SKUOption{SKUId: sku.ID, OptionID: opt.ID}
      err = s.skuRepo.SaveSKUOption(&skuOpt)
      if err != nil {
          return nil, err
      }
  }

  // Save the price to the pricing APIServer
  req := SetRequest{
      Code:         sku.Code,
      Name:         product.Name,
      BasePrice:    math.Round(price * 100.0),
      CurrencyCode: "ZAR",
  }

  _, err = s.pricingClient.CreatePrice(req, token)

  return &sku, err
}

func  (s *SkuService) GetSKUPrice(sku models.SKU, token string) (models.SKU, error) {
  // This also needs a cache!
  identity, err := s.identityClient.GetContext(token)
  if err != nil {
    // Visitors would not be able to get a price, so ofer them some defaults
    identity.RegionCode = "GP"
    identity.CustomerSegment = "VISITOR"
    identity.Currency = "ZAR"
  }

  var response models.SKU

  // Check cache first
  cached, _ := s.cacheRepo.FindValidPrice(sku.Code, identity.RegionCode, identity.CustomerSegment, identity.Currency)
  if cached != nil {
    response = models.SKU{
      ID:             sku.ID,
      ProductID:      sku.ProductID,
      Code:           sku.Code,
      Price:          cached.OriginalPrice,
      Stock:          0, // Inventory is Oos for now
    }
  } else {
      // Call pricing service if not cached
      req := PricingRequest{
          SkuCode:         sku.Code,
          Quantity:        1,
          RegionCode:      identity.RegionCode,
          CustomerSegment: identity.CustomerSegment,
          TargetCurrency:  identity.Currency,
      }
      resp, err := s.pricingClient.Calculate(req, token)
      if err != nil {
          // Should test for 404.  If SKU not found in pricing service either, then
          // return a negative price.
          return models.SKU{
            ID:             sku.ID,
            ProductID:      sku.ProductID,
            Code:           sku.Code,
            Price:          -math.MaxInt,
            Stock:          0, // Inventory is OoS for now
          }, nil  //NOTE not err!  I dont like this code.
      }

      // Save to cache
      s.cacheRepo.SavePrice(&models.PriceCache{
          CacheID:        uuid.NewString(),
          SkuCode:        sku.Code,
          RegionCode:     identity.RegionCode,
          CustomerSegment: identity.CustomerSegment,
          CurrencyCode:   identity.Currency,
          OriginalPrice:  resp.OriginalPrice,
          FinalPrice:     resp.FinalPrice,
          AppliedRules:   utils.SerializeRules(resp.AppliedRules),
          CachedAt:       time.Now(),
          ExpiresAt:      time.Now().Add(15 * time.Minute),
      })

      response = models.SKU{
        ID:             sku.ID,
        ProductID:      sku.ProductID,
        Code:           sku.Code,
        Price:          resp.OriginalPrice,
        Stock:          0, // Inventory is OoS for now
      }
  }
  return response, err
}

func (s *SkuService) GetSKU(ctx context.Context, ID string, token string) (models.SKU, error) {
    sku, err := s.skuRepo.FindByID(ID)
    if err != nil {
        return models.SKU{}, err
    }

    return s.GetSKUPrice(*sku, token)
}

func (s *SkuService) GetProductSKUList(ctx context.Context, productID string, token string) ([]dto.SKUListItem, error) {
  SKUList, err := s.skuRepo.GetProductSKUList(productID)
  if err != nil {
      return []dto.SKUListItem{}, err
  }

  var resp []dto.SKUListItem
  for _,sku := range SKUList {
    sk, err := s.GetSKUPrice(sku, token)
    if err!=nil {
      continue
    }   // Dont return a SKU if it has no price!  Test the use case!!
    resp = append(resp, dto.SKUListItem {
      ID   : sk.ID,
      Code : sk.Code,
      Price: sk.Price,
      Stock: sk.Stock,
      })
  }
  return resp, err
}

/*  skuList, err := s.skuRepo.GetProductSKU(id)
  if err != nil {
      return []dto.SKUListItem{}, err
  }

  for _, sku := range skuList {
    response, err := s.GetSKUPrice(sku, token)
    if err!=nil { continue }
    sku.Price := response.Price
  }

  return skuList, err
}
*/
