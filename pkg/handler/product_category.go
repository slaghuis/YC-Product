package handler

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func (h *Handler) UpdateProductCategories(c *gin.Context) {
  var req struct {
    ProductID   	uint		`json:"product_id"`
    CategoryIDs  	[]uint	`json:"categories"`
  }
  if err := c.BindJSON(&req); err != nil {
      c.AbortWithError(http.StatusBadRequest, err)
      return
  }

  err := h.ProductService.UpdateProductCategories(c, req.ProductID, req.CategoryIDs)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Categories not updated"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"status":"product categores updated"})
}


func (h *Handler) ListProductCategories(c *gin.Context) {
  id := c.Param("id")  //

  categoryResponse, err := h.ProductService.ReadProductCategories(c, id)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, categoryResponse)
}
