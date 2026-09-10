package repositories

import (
  "github.com/slaghuis/YC-Product/pkg/dto"
  "github.com/slaghuis/YC-Product/pkg/models"
)

func (r *ProductRepository)  ReadProductParameters(id string) ([]dto.ProductParameterItem, error) {
  var items []dto.ProductParameterItem

  err := r.db.Table("categories AS parent").
      Select(`child.name AS parameter_name,
              child.id AS category_id,
              parameters.id AS parameter_id,
              parameters.markdown AS markdown,
              parameters.product_id AS product_id`).
              Joins("JOIN categories AS child ON child.parent_id = parent.id").
              Joins("LEFT JOIN product_parameters AS parameters ON child.id = parameters.category_id AND parameters.product_id = ?", id).
              Where("parent.alternate_key = ?", "parameters").
              Scan(&items).Error
  return items, err
}


func (r *ProductRepository) SaveProductParameters(productParameters []models.ProductParameter) (error){
  return r.db.Save(productParameters).Error
}
