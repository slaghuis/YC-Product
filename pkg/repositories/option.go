package repositories

import (
    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
)

func (r *ProductRepository) CreateOptions(options []models.Option) error {
  return r.db.Create(&options).Error
}

func (r *ProductRepository) GetOption(id uint) (*models.Option, error) {
  var item models.Option
  err := r.db.First(&item, id).Error
  return &item, err
}

func (r *ProductRepository) GetOptions(OptionIDs []uint) ([]models.Option, error) {
  var options []models.Option
  err := r.db.Where("id IN ?", OptionIDs).Find(&options).Error
  return options, err
}

func (r *ProductRepository) GetOptionItemsByVariant(id uint) ([]dto.OptionItem, error) {

  var optionItems []dto.OptionItem
  if err := r.db.Table("options").
    Select(`
        options.id AS option_id,
        options.value
    `).
    Joins("JOIN variants on variants.id=options.variant_id").
    Where("variants.id = ?", id).
    Order("options.id ASC").
    Scan(&optionItems).Error; err != nil {
      return []dto.OptionItem{}, err
  }

  return optionItems, nil
}

func (r *ProductRepository) SaveOption(item *models.Option) error {
  return r.db.Save(item).Error
}

func (r *ProductRepository) DeleteOption(optionID uint) error {
  return r.db.Delete(&models.Option{}, optionID).Error
}

func (r *ProductRepository) GetOptionItems(variantID uint) ([]dto.OptionItem, error) {
  var options []dto.OptionItem
  err := r.db.Table("options").
    Select(`
        options.id AS option_id,
        options.value
        `).
    Joins("JOIN variants on variants.id=options.variant_id").
    Where("variants.id = ?", variantID).
    Order("options.id ASC").
    Scan(&options).Error
  return options, err
}

func (r *ProductRepository) UpdateOption(optionID uint, value string) (models.Option, error) {
  var item models.Option
  err := r.db.First(&item, optionID).Error
  if err != nil {
    return models.Option{}, err
  }

  item.Value = value

  err =r.db.Save(&item).Error

  return item, err
}
