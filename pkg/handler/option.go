package handler
/* ************************************************************************************
 * Usage Flow
 * ************************************************************************************ */

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// Responds to auth.GET("/variants/:id", h.GetVariantOptions)  //get one variant and its options
func (h *Handler) GetVariantOptions(c *gin.Context) {
  id := c.Param("id")

  variant, err := h.ProductService.GetVariant(c, id)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Option not found"})
    return
  }

  c.JSON(http.StatusOK, variant)
}

/*  var item models.Variant
  if err := h.db.
      First(&item, id).Error; err != nil {
      c.JSON(http.StatusNotFound, gin.H{"error": "Option not found"})
      return
  }

  var options []OptionItem
  if err := h.db.Table("options").
      Select(`
          options.id AS option_id,
          options.value
      `).
      Joins("JOIN variants on variants.id=options.variant_id").
      Where("variants.id = ?", id).
      Order("options.id ASC").
      Scan(&options).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, gin.H{
    "id"         : item.ID,
    "name"       : item.Name,
    "product_id" : item.ProductID,
    "options" 	 : options,
  })
}
*/

// GET a single option
// Responds to auth.GET("/options/:id", h.GetOption)
func (h *Handler) GetOption(c *gin.Context) {
    id := c.Param("id")
    item, err := h.ProductService.GetOption(c, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Option not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
      "id":         item.ID,
      "variant_id": item.VariantID,
      "value":      item.Value,
    })
}

func (h *Handler) UpdateOption(c *gin.Context) {
  id := c.Param("id")
  var req struct {
    Value       string  `json:"value"`
  }
  if err := c.ShouldBindJSON(&req); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  item, err := h.ProductService.UpdateOption(c, id, req.Value)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()} )
      return
  }

  c.JSON(http.StatusOK, &item)
}


func (h *Handler) CreateOption(c *gin.Context) {
  id := c.Param("id")  // variantID
  var req struct {
    ID            uint        `json:"id"`
    Options       []string    `json:"options"`
  }
  if err := c.ShouldBindJSON(&req); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  if err := h.ProductService.CreateOptions(c, id, req.Options); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  // Now load and assemble the variant detail to RETURN
  item , err := h.ProductService.GetVariant(c, id)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, item)   // item of type dto.Variant
/*

  c.JSON(http.StatusOK, gin.H{
    "id"         : item.ID,
    "name"       : item.Name,
    "product_id" : item.ProductID,
    "options" 	 : persistedOptions,
  })
*/
}

func (h *Handler) DeleteOption(c *gin.Context) {
  id := c.Param("id")

  err := h.ProductService.DeleteOption(c, id)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "message": "option successfully deleted",
  })
}
