package services

import (
    "context"

    "github.com/slaghuis/YC-Product/pkg/models"
)

func (s *productService) CreateOptions(ctx context.Context, id string, opt []string) (error) {

  variantID, err := StrToUint(id)
  if err != nil {
    return  err
  }

  var options []models.Option
  for _, op := range opt{
    options = append(options, models.Option{
      VariantID: variantID,
      Value:     op,
    })
  }

  return s.productRepo.CreateOptions(options)
}


func (s *productService) GetOption(ctx context.Context, optionID string) (*models.Option, error) {
  id, err := StrToUint(optionID)
  if err != nil {
    return nil, err
  }
  return s.productRepo.GetOption(id)
}

func (s *productService) UpdateOption(ctx context.Context, optionID, value string) (models.Option, error) {
  id, err := StrToUint(optionID)
  if err != nil {
    return models.Option{}, err
  }

  return s.productRepo.UpdateOption(id, value)
}

func (s *productService) DeleteOption(ctx context.Context, optionID string) (error) {
  id, err := StrToUint(optionID)
  if err != nil {
    return err
  }
  return s.productRepo.DeleteOption(id)
}
