package handler
/* ************************************************************************************
 * Usage Flow
 * ************************************************************************************ */

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategory(c *gin.Context) {
    var req struct {
      AlternateKey  string   `json:"alternate_key"`
      ParentId      uint     `json:"parent_id"`
      Name          string   `json:"name"`
      IconPath      string   `json:"icon_path"`
      SubType       uint     `json:"sub_type"`
      Description   string   `json:"description"`
      SortOrder     uint     `json:"sort_order"`
    }

    if err := c.BindJSON(&req); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    item, err := h.ProductService.CreateCategory(c.Request.Context(), req.Name, req.AlternateKey, req.IconPath, req.Description, req.ParentId, req.SubType, req.SortOrder)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "category not created"})
        return
    }

    c.JSON(http.StatusOK, &item)
}


func (h *Handler) UpdateCategory(c *gin.Context) {
    id := c.Param("id")   // Do I use this one or lift it off the request body

    var req struct {
      ID            uint     `json:"id"`
      AlternateKey  string   `json:"alternate_key"`
      ParentId      uint     `json:"parent_id"`
      Name          string   `json:"name"`
      IconPath      string   `json:"icon_path"`
      SubType       uint     `json:"sub_type"`
      Description   string   `json:"description"`
      SortOrder     uint     `json:"sort_order"`
    }

    if err := c.BindJSON(&req); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    item, err := h.ProductService.UpdateCategory(c.Request.Context(), id, req.Name, req.AlternateKey, req.IconPath, req.Description, req.ParentId, req.SubType, req.SortOrder)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
        return
    }

    c.JSON(http.StatusOK, &item)
}

func (h *Handler) GetCategory(c *gin.Context) {
    id := c.Param("id")
    item, err := h.ProductService.GetCategory(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
      "id"            : item.ID,
      "alternate_key" : item.AlternateKey,
      "parent_id"     : item.ParentId,
      "name"          : item.Name,
      "icon_path"     : item.IconPath,
      "sub_type"      : item.SubType,
      "description"   : item.Description,
      "sort_order"    : item.SortOrder,
    })
}

/* DANGEROUS - COULD RETURN A LOT - Not sure this is even used at this stage */
func (h *Handler) ListCategories(c *gin.Context) {
  items, err := h.ProductService.ListCategories(c.Request.Context())
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusOK,  gin.H{
    "categories": items,
  })
}


/* /products/categories/parent?alt=top_tier    *
 * /products/categories/parent?id=64           *
 * Tested 20 Feb 2026                          */

func (h *Handler) ListCategoriesByParent(c *gin.Context)  {
  id := c.DefaultQuery("id", "nan")
  alt := c.DefaultQuery("alt", "")

  items, err := h.ProductService.GetCategoriesByParent(c.Request.Context(), id, alt)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  //h.LogJson(items)

  c.JSON(http.StatusOK,  gin.H{
  	"categories": items,
  })

}

func (h *Handler) ListMerchantCategoriesByParent(c *gin.Context)  {
  id := c.DefaultQuery("id", "nan")
  alt := c.DefaultQuery("alt", "")

  items, err := h.ProductService.GetMerchantCategoriesByParent(c.Request.Context(), id, alt)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

//  c.JSON(http.StatusOK,items)
  c.JSON(http.StatusOK,  gin.H{
  	"rows": items,
  })
}

func (h *Handler) DeleteCategory(c *gin.Context) {
  id := c.Param("id")
  err := h.ProductService.DeleteCategory(c.Request.Context(), id)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }
  c.JSON(http.StatusOK, gin.H{
    "status"  : "success",
    "message" : "category deleted",
  })
}

func (h *Handler) GetCategoryBreadcrumb(c *gin.Context) {
  categoryId := c.Param("id")
  if !h.IsNumeric(categoryId) {
      c.JSON(http.StatusBadRequest,  gin.H{"error": "parameter is not numeric"})
  }

  breadcrumbs, err := h.ProductService.GetCategoryBreadcrumb(c.Request.Context(), categoryId)
  if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
  }

  c.JSON(http.StatusOK,  gin.H{
  	"rows": breadcrumbs,
  })
}
