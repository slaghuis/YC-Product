package services

import (
    "context"
    "errors"

    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
//    "github.com/slaghuis/YC-Product/pkg/repositories"
)

func (s *productService) CreateCategory(ctx context.Context, name, alternateKey, iconPath, description string, parentId, subType, sortOrder uint) (*models.Category, error) {
  item := &models.Category {
    // Overwrite all except the ID
    AlternateKey : alternateKey,
    ParentId     : parentId,
    Name         : name,
    IconPath     : iconPath,
    SubType      : subType,
    Description  : description,
    SortOrder    : sortOrder,
  }

  err := s.productRepo.SaveCategory(item)
  return item, err
}

func (s *productService) UpdateCategory(ctx context.Context, id, name, alternateKey, iconPath, description string, parentId, subType, sortOrder uint) (*models.Category, error) {
  categoryId, err := StrToUint(id)
  if err != nil {return nil, err}

  item, err := s.productRepo.GetCategory(categoryId)
  if err != nil {return nil, err}

  // Overwrite all except the ID
  item.AlternateKey = alternateKey
  item.ParentId     = parentId
  item.Name         = name
  item.IconPath     = iconPath
  item.SubType      = subType
  item.Description  = description
  item.SortOrder    = sortOrder

  err = s.productRepo.SaveCategory(item)
  return item, err
}

func (s *productService) GetCategoryBreadcrumb(ctx context.Context, id string) ([]models.Breadcrumb, error) {
  categoryId, err := StrToUint(id)
  if err != nil {
    return nil, err
  }
  return s.productRepo.GetCategoryBreadcrumb(categoryId)
}

func (s *productService) DeleteCategory(ctx context.Context, id string) error {
  categoryId, err := StrToUint(id)
  if err != nil {
    return err
  }
  return s.productRepo.DeleteCategory(categoryId)
}

func (s *productService) GetMerchantCategoriesByParent(ctx context.Context, id, alt string) ([]dto.CategoryMerchantListItem, error) {
  var items []dto.CategoryMerchantListItem
  parentId, err := StrToUint(id)
  if err != nil {
    if len(alt)>0 {
      items, err = s.productRepo.ListMerchantCategoriesByAltKey(alt)
    } else {
      err = errors.New("bad request.  need id=x or alt=.. ")
    }
  } else {
    items, err = s.productRepo.ListMerchantCategoriesByParentId(parentId)
  }
  return items, err
}

func (s *productService) GetCategoriesByParent(ctx context.Context, id, alt string) ([]dto.CategoryListItem, error) {
  var items []dto.CategoryListItem
  parentId, err := StrToUint(id)
  if err != nil {
    if len(alt)>0 {
      items, err = s.productRepo.ListCategoriesByAltKey(alt)
    } else {
      err = errors.New("bad request.  need id=x or alt=.. ")
    }
  } else {
    items, err = s.productRepo.ListCategoriesByParentId(parentId)
  }
  return items, err
}

func (s *productService) ListCategories(ctx context.Context) ([]dto.CategoryListItem, error) {
  return s.productRepo.ListCategories()
}

func (s *productService) ListProductCategoryItem(ctx context.Context, alternateKey, id string) ([]dto.ProductCategoryItem, error) {
  return s.productRepo.ListProductCategoryItem(alternateKey, id)
}

func (s *productService) GetCategory(ctx context.Context, id string) (*models.Category, error) {
  categoryID, err := StrToUint(id)
  if err != nil { return nil, err }

  return s.productRepo.GetCategory(categoryID)
}
