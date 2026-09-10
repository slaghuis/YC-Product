package repositories

import (
    "gorm.io/gorm"
    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
)

type ReviewRepository struct {
    db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
    return &ReviewRepository{db: db}
}

func (r *ReviewRepository) SaveReview(item *models.Review) error {
  return r.db.Save(item).Error
}

func (r *ReviewRepository) DeleteReview(id uint) (error) {
  return r.db.Delete(&models.Review{}, id).Error
}

func (r *ReviewRepository) GetReview(id uint) (*models.Review, error) {
  var item models.Review
  err := r.db.First(&item, id).Error
  return &item, err
}

func (r *ReviewRepository) ListReviews(userID string) ([]dto.ReviewListItem, error) {
  var items []dto.ReviewListItem
  err := r.db.Model(&models.Review{}).
        Select("reviews.score, reviews.comment, reviews.moderated, reviews.full_name, reviews.account_id, products.name as product_name, reviews.product_id").
        Joins("left join products on reviews.product_id=reviews.product_id").
        Where("account_id = ?", userID).
        Scan(&items).
        Order("reviews.updated_at desc").Error
  return items, err
}

func (r *ReviewRepository) ListReviewsByProduct(productID string) ([]dto.ReviewListItem, error) {
  var items []dto.ReviewListItem
  err := r.db.Model(&models.Review{}).
        Select("reviews.score, reviews.comment, reviews.moderated, reviews.full_name, reviews.account_id, products.name as product_name, reviews.product_id").
        Joins("left join products on reviews.product_id=reviews.product_id").
        Where("products.id = ?", productID).
        Scan(&items).
        Order("reviews.updated_at desc").Error
  return items, err
}

func (r *ReviewRepository) ToggleLike(productID uint, accountID string) (bool, error) {
  var item models.Like

  if result := r.db.First(&item).Where("product_id = ?", productID).Where("account_id = ?", accountID); result.Error != nil {
    item.Score=0;
    item.ProductID = productID
    item.AccountID = accountID
  }

  // Do the toggle with some simple math
  item.Score       = 1-item.Score

  err := r.db.Save(&item).Error
  return (item.Score==1), err
}


func (r *ReviewRepository) CountLikes(productID string) (int64, error) {
    var totalSum int64
    err := r.db.Table("likes").
        Select("sum(score)").
        Where("product_id = ?", productID).
        Scan(&totalSum).Error
    return totalSum, err
}


/* GOOD CODE, BUT NOT USED
func (r *ReviewRepository) ListByCustomerID(customerID string) ([]models.Review, error) {
  var list []models.Review
  err := r.db.Where("account_id = ?", customerID).Find(&list).Error
  return list, err
}

func (r *ReviewRepository) ListByProductID(productID string) ([]models.Review, error) {
  var list []models.Review
  err := r.db.Where("product_id = ?", productID).Find(&list).Error
  return list, err
}

*/
