package handler
/* ************************************************************************************
 * Usage Flow
 * ************************************************************************************ */

import (
  "net/http"
  "github.com/gin-gonic/gin"
//	"github.com/slaghuis/YC-Product/pkg/models"
)

func (h *Handler) CreateSupplier(c *gin.Context) {
  var req struct {
    Name          string `json:"name"`
    RepName       string `json:"rep_name"`
    Phone         string `json:"phone"`
//    PhoneVerified bool   `json:"phone_verified"`
    Email         string `json:"email"`
  }
  if err := c.ShouldBindJSON(&req); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  supplier, err := h.SupplierService.CreateSupplier(c, req.Name, req.RepName, req.Phone, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

  c.JSON(http.StatusCreated, &supplier)
}

func (h *Handler) UpdateSupplier(c *gin.Context) {
  id := c.Param("id")
  var req struct {
    ID            uint   `json:"id"`
    Name          string `json:"name"`
    RepName       string `json:"rep_name"`
    LogoPath      string `json:"logo_path"`
    Phone         string `json:"phone"`
    PhoneVerified bool   `json:"phone_verified"`
    Email         string `json:"email"`
    SMS           bool   `json:"sms"`
    WhatsApp      bool   `json:"whatsapp"`
    EmailNotify   bool   `json:"email_notify"`
    AppPush       bool   `json:"app_push"`
  }

  if err := c.BindJSON(&req); err != nil {
      c.AbortWithError(http.StatusBadRequest, err)
      return
  }

  supplier, err := h.SupplierService.UpdateSupplier(c, id, req.Name, req.RepName, req.LogoPath, req.Phone, req.PhoneVerified, req.Email, req.SMS, req.WhatsApp, req.EmailNotify, req.AppPush)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

  c.JSON(http.StatusCreated, &supplier)
}

func (h *Handler) GetSupplier(c *gin.Context) {
    id := c.Param("id")

    item, err := h.SupplierService.GetSupplierBasic(c.Request.Context(), id)
  	if err != nil {
  		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
  		return
  	}

    c.JSON(http.StatusOK, item)
}

func (h *Handler) ListSuppliers(c *gin.Context) {
  token := c.GetString("token_key")

  suppliers, err := h.SupplierService.GetBasicSupplierList(c.Request.Context(), token)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }
  c.JSON(http.StatusOK, gin.H{"suppliers": suppliers})
}

func (h *Handler) DeleteSupplier(c *gin.Context) {
  id := c.Param("id")  // Variant_id in this case
  err := h.SupplierService.DeleteSupplier(c.Request.Context(), id)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, gin.H{"status": "supplier deleted"})
}

/* ************************************************************ *
 * Lists all products with a merchant summary
 * Typycally used in the product administration function
 * Parameters:
 *   limit=x Page size
 *   page=y  Page number
 *   sort=z  Sort Order
 * *********************************************************** */
func (h *Handler) ListMerchantSuppliers(c *gin.Context) {
  page := c.DefaultQuery("page", "1")
  limit := c.DefaultQuery("limit", "20")
  sort := c.DefaultQuery("sort", "id desc")

  pagination, err := h.SupplierService.GetMerchantProductSupliers(page, limit, sort)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, pagination)
}
