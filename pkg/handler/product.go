package handler
/* ************************************************************************************
 * Usage Flow
 * ************************************************************************************ */

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"github.com/slaghuis/YC-Product/pkg/models"
)

func (h *Handler) CreateProduct(c *gin.Context) {
  var req struct {
    Name          string   			 `json:"name"`
  	Description 	string         `json:"description"`
  	SupplierID  	uint		 			 `json:"supplier_id"`
  }

  if err := c.BindJSON(&req); err != nil {
      c.AbortWithError(http.StatusBadRequest, err)
      return
  }

  item, err := h.ProductService.CreateProduct(c.Request.Context(), req.Name, req.Description, req.SupplierID)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusCreated, &item)
}

func (h *Handler) GetProductBasicInfo(c *gin.Context) {
    id := c.Param("id")

    item, err := h.ProductService.GetProductBasicInfo(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, &item)
}

/* ************************************************************ *
 * Lists all products
 * Typycally used in the product administration function
 * Parameters:
 *   limit=x Page size
 *   page=y  Page number
 *   sort=z  Sort Order
 *   category=n where the product belong to categtories.id=n
 * *********************************************************** */
func (h *Handler) ListProducts(c *gin.Context) {
  page := c.DefaultQuery("page", "1")
  limit := c.DefaultQuery("limit", "20")
  sort := c.DefaultQuery("sort", "relevance")
  category := c.DefaultQuery("category", "")

  token := c.GetString("token_key")

  pagination, err := h.ProductService.ListProducts(c.Request.Context(), page, limit, sort, category, token)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

	c.JSON(http.StatusOK, pagination)
}

// Count products that fulfill the current product preferences
func (h *Handler) CountProducts(c *gin.Context) {
  category := c.DefaultQuery("category", "")

  token := c.GetString("token_key")

  count, err := h.ProductService.CountProducts(c.Request.Context(), category, token)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

	c.JSON(http.StatusOK, count)
}

/* ************************************************************ *
 * Lists all products with the intent to inform the merchant
 * of what is offered for sale.
 * Parameters:
 *   limit=x Page size
 *   page=y  Page number
 *   sort=z  Sort Order
 * *********************************************************** */
 func (h *Handler) ListMerchantProducts(c *gin.Context) {
   page := c.DefaultQuery("page", "1")
   limit := c.DefaultQuery("limit", "20")
   sort := c.DefaultQuery("sort", "id desc")

   pagination, err := h.ProductService.ListMerchantProducts(c.Request.Context(), page, limit, sort)
   if err != nil {
     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
     return
   }

 	c.JSON(http.StatusOK, pagination)
 }

func (h *Handler) UpdateProductDefinition(c *gin.Context) {
//  body, _ := io.ReadAll(c.Request.Body)
//  fmt.Println("Raw body:", string(body))
//  c.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // reset for BindJSON


    id := c.Param("id")   // Do I use this one or lift it off the request body

    var req struct {
      Name          string   			          `json:"name"`
      Description 	string                  `json:"description"`
      State         models.ProductState    	`json:"state"`
    }
    if err := c.BindJSON(&req); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }
    item, err := h.ProductService.UpdateProduct(c.Request.Context(), id, req.Name, req.Description, req.State)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, &item)
}

func (h *Handler) SetState(c *gin.Context) {
  id := c.Param("id")  // Product_id in this case
  state := c.Query("state")

  err := h.ProductService.SetState(c.Request.Context(), id, state)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
    "message": "Product state changed successfully",
  })
}
