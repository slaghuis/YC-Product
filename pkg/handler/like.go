package handler
/* ************************************************************************************
 * Usage Flow
 * ************************************************************************************ */

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func (h *Handler) ToggleLike(c *gin.Context) {
  productID := c.Param("id")  // Product_id in this case
  accountID := c.GetString("user_id")
  like, err := h.ReviewService.ToggleLike(c, productID, accountID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "like": like,
    })
}

func (h *Handler) CountLikes(c *gin.Context) {
  id := c.Param("id")
  count, err := h.ReviewService.CountLikes(c, id)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "count": count,
  })
}


/* DONT THINK THIS IS USED

func (h *Handler) GetLike(c *gin.Context) {
    id := c.Param("id")
    var item models.Like

    var Score uint = 0

    if result := h.db.First(&item, id).Where("product_id = ?", id).Where("account_id = ?", c.GetString("user_id")); result.Error == nil {
      Score = item.Score
    }

    c.JSON(http.StatusOK, gin.H{"score": Score})
}
*/
