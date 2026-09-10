package handler
/* ************************************************************************************
 * Usage Flow
 * ************************************************************************************ */

import (
  "net/http"
  "github.com/gin-gonic/gin"
//	"github.com/slaghuis/YC-Product/pkg/dto"
)

/* Represents a dimension of variation (e.g., "Color", "Size"). */
func (h *Handler) CreateVariant(c *gin.Context) {
  var req struct {
    Name      string      `json:"name"` // e.g., "Color"
    Options   []string    `json:"options"`
  }
  if err := c.BindJSON(&req); err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }

  id := c.Param("id")  // Product_id in this case
  variant, err := h.ProductService.CreateVariant(c, id, req.Name, req.Options)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"variant": variant})
}

func (h *Handler) GetVariantList(c *gin.Context) {
  id := c.Param("id")  // Product_id in this case

  productId, err := h.StrToUint(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  variantGroups, err := h.ProductService.GetVariantList(c, id)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "id": productId,
    "variants":	variantGroups,
  })
}

func (h *Handler) DeleteVariant(c *gin.Context) {
  id := c.Param("id")  // Variant_id in this case
  err := h.ProductService.DeleteVariant(c, id)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, gin.H{"status": "variant deleted"})
}
