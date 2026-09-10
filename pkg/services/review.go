package services

import (
    "context"

    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
    "github.com/slaghuis/YC-Product/pkg/repositories"
)

type reviewService struct {
    repo           *repositories.ReviewRepository
}

func NewReviewService(repo *repositories.ReviewRepository) ReviewService {
    return &reviewService{
      repo: repo,
    }
}

func (s *reviewService) CreateReview(ctx context.Context, score uint, comment, name, accountID string, productID uint) (*models.Review, error) {
  item := &models.Review {
    Score     : score,
    Comment   : comment,
    Moderated : false,
    FullName  : name,
    AccountID : accountID,
    ProductID : productID,
  }
  err := s.repo.SaveReview(item)
  return item, err
}

func (s *reviewService) UpdateReview(ctx context.Context, id_str, comment string, score uint, moderated bool) (*models.Review, error) {
  id, err := StrToUint(id_str)
  if err != nil {
    return nil, err
  }

  item := &models.Review {
    ID        : id,
    Score     : score,
    Comment   : comment,
    Moderated : moderated,
  }
  err = s.repo.SaveReview(item)
  return item, err
}

func (s *reviewService) GetReview(ctx context.Context, id_str string) (*models.Review, error) {
  id, err := StrToUint(id_str)
  if err != nil {
    return &models.Review{}, nil
  }
  return s.repo.GetReview(id)
}

func (s *reviewService) ListReviews(ctx context.Context, userID string) ([]dto.ReviewListItem, error) {
  return s.repo.ListReviews(userID)
}

func (s *reviewService) ListReviewsByProduct(ctx context.Context, productID string) ([]dto.ReviewListItem, error) {
  return s.repo.ListReviewsByProduct(productID)
}

func (s *reviewService) DeleteReview(ctx context.Context, reviewID string) (error) {
  id, err := StrToUint(reviewID)
  if err != nil {
    return err
  }
  return s.repo.DeleteReview(id)
}
