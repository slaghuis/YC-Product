package handler
/* ************************************************************************************
 * Usage Flow
 * ************************************************************************************ */

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"github.com/slaghuis/YC-Product/pkg/dto"
  "github.com/slaghuis/YC-Product/pkg/models"
)

func (h *Handler) ListProductParameters(c *gin.Context) {
  id := c.Param("id")  // Product_id in this case

  items, err := h.ProductService.ReadProductParameters(c, id)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, items)
}


func (h *Handler) UpdateProductParameters(c *gin.Context) {
  id := c.Param("id")
  productID, err := h.StrToUint(id)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "message": "failed to read product id: " + err.Error(),
    })
    return
  }

  var request_body []dto.ParameterListItem
  if err := c.BindJSON(&request_body); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "message": "failed to parse request body: " + err.Error(),
    })
    return
  }

//  h.LogJson(request_body)

  productParameters := []models.ProductParameter{}
  //Loop possible parameters
  for _, item := range request_body {
    new_product_parameter := models.ProductParameter{
      ID:         item.ParameterId,
      Markdown:   item.Markdown,
      CategoryID: item.CategoryId,
      ProductID:  productID,
    }
    productParameters = append(productParameters, new_product_parameter)
  }

  if err := h.ProductService.SaveProductParameters(c, productParameters); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "status": "error",
      "message": "failed to update product parameters: " + err.Error(),
    })
    return
  }

/*
  items, err := h.ProductService.ReadProductParameters(c, id)

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "status": "error",
      "message": "failed to read product parameters: " + err.Error(),
    })
    return
  }
*/
  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "message": "product parameters saved",
  })
}
