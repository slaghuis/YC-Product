package handler
/* ************************************************************************************
 * Usage Flow
 * ************************************************************************************ */

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"github.com/slaghuis/YC-Product/pkg/models"
)

func (h *Handler) CreateReview(c *gin.Context) {
  var req struct {
    Score       uint            `json:"score"`
    Comment     string          `json:"comment"`
    ProductID   uint            `json:"product_id"`
  }
  if err := c.BindJSON(&req); err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }

  name := c.GetString("user_name")
  accountID := c.GetString("user_id")
  review, err := h.ReviewService.CreateReview(c, req.Score, req.Comment, name, accountID, req.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

  c.JSON(http.StatusCreated, &review)
}

func (h *Handler) UpdateReview(c *gin.Context) {
  var req struct {
    Score       uint            `json:"score"`
    Comment     string          `json:"comment"`
    Moderated   bool            `json:"moderated"`
  }
  if err := c.BindJSON(&req); err != nil {
      c.AbortWithError(http.StatusBadRequest, err)
      return
  }

  id := c.Param("id")
  review, err := h.ReviewService.UpdateReview(c, id, req.Comment, req.Score, req.Moderated)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, &review)
}

func (h *Handler) GetReview(c *gin.Context) {
    id := c.Param("id")
    item, err := h.ReviewService.GetReview(c, id)
    if err != nil {
      c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
      return
    }
    c.JSON(http.StatusOK, item)
}

func (h *Handler) ListReviews(c *gin.Context) {
  userID := c.GetString("user_id")
  reviews, err := h.ReviewService.ListReviews(c, userID)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }
  c.JSON(http.StatusOK, gin.H{
  	"reviews": reviews,
  })
}

func (h *Handler) ListProductReviews(c *gin.Context) {
  id := c.Param("id")  // Product_id in this case
  reviews, err := h.ReviewService.ListReviewsByProduct(c, id)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }
  c.JSON(http.StatusOK, gin.H{
  	"reviews": reviews,
  })
}


func (h *Handler) DeleteReview(c *gin.Context) {
  id := c.Param("id")

  var item models.Review

  if result := h.db.First(&item, id); result.Error != nil {
      c.AbortWithError(http.StatusNotFound, result.Error)
      return
  }

  if result := h.db.Delete(&item); result.Error != nil {
      c.AbortWithError(http.StatusInternalServerError, result.Error)
      return
  }

  c.Status(http.StatusOK)  //c.JSON
}
