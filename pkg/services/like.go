package services

//Should be removed from this service.  Has been moved to preference service

import (
    "context"
)

func (s *reviewService) ToggleLike(ctx context.Context, pid, accountID string) (bool, error) {
  productID, err := StrToUint(pid)
  if err != nil {
    return false, err
  }

  return s.repo.ToggleLike(productID, accountID)
}

func (s *reviewService) CountLikes(ctx context.Context, productID string) (int64, error) {
  return s.repo.CountLikes(productID)
}
