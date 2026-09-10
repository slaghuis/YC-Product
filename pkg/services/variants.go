package services

import (
    "context"
    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
)

func (s *productService) CreateVariant(ctx context.Context, id, name string, options []string) (*models.Variant, error) {

  productID, err := StrToUint(id)
  if err != nil {
    return nil, err
  }

  item := &models.Variant {
    ProductID:     productID,
    Name:          name,
  }
  err = s.productRepo.SaveVariant(item)
  if err != nil {
    return nil, err
  }

  var optionList []models.Option
  for _, opt := range options{
    optionList = append(optionList, models.Option{
      VariantID: item.ID,
      Value:     opt,
    })
  }

  err = s.productRepo.CreateOptions(optionList)
  return item, err
}


func (s *productService) GetVariant(ctx context.Context, id string) (*dto.Variant, error) {
  variantID, err := StrToUint(id)
  if err != nil {
    return nil, err
  }

  variant, err := s.productRepo.GetVariant(variantID)
  if err != nil {
    return nil, err
  }

  optionItems, err := s.productRepo.GetOptionItemsByVariant(variant.ID)
  if err != nil {
    return nil, err
  }

  return &dto.Variant{
  	  VariantID:  variant.ID,
  	  Name:       variant.Name,
  	  ProductID:  variant.ProductID,
  	  Options:    optionItems,
  }, nil

}

func (s *productService) GetVariantList(ctx context.Context, id string) ([]dto.VariantGroup, error) {

  variants, err := s.productRepo.GetVariantList(id)
  if err != nil {
    return nil, err
  }

  VariantGroup := make([]dto.VariantGroup, 0, len(variants))

  // Loop through variants and run second query for each to read the options
  for _, bk := range variants {

    options, err := s.productRepo.GetOptionItems(bk.ID)
    if err != nil {
      continue
    }
    // Add group with categories
    VariantGroup = append(VariantGroup, dto.VariantGroup{
      VariantId: 		bk.ID,
      VariantName:  bk.Name,
      Options:   		options,
    })
  }
  return VariantGroup, err
}

func (s *productService) DeleteVariant(ctx context.Context, id string) error {
  variantId, err := StrToUint(id)
  if err != nil {
    return err
  }

  return s.productRepo.DeleteVariant(variantId)
}
