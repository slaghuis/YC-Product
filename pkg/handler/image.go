package handler
/* ************************************************************************************
 * Usage Flow
 * ************************************************************************************ */

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"github.com/slaghuis/YC-Product/pkg/models"
)

type AddImageRequestBody struct {
    Name          string `json:"name"`
    AltText       string `json:"alt_text"`
    ImagePath     string `json:"image_path"`
    Type          models.ImageType   `json:"type"`
    ProductID     uint `json:"product_id"`
}

func (h *Handler) CreateImage(c *gin.Context) {
    body := AddImageRequestBody{}

    if err := c.BindJSON(&body); err != nil {
          c.AbortWithError(http.StatusBadRequest, err)
          return
    }

    var item models.Image
    item.Name          = body.Name
    item.AltText       = body.AltText
    item.ImagePath     = body.ImagePath
    item.Type          = body.Type
    item.ProductID     = body.ProductID

    if result := h.db.Create(&item); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &item)
}

type UpdateImageRequestBody struct {
    Name          string `json:"name"`
    AltText       string `json:"alt_text"`
    ImagePath     string `json:"image_path"`
    Type          models.ImageType   `json:"type"`
}

func (h *Handler) UpdateImage(c *gin.Context) {
    id := c.Param("id")
    body := UpdateImageRequestBody{}

    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var item models.Image

    if result := h.db.First(&item, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    item.Name          = body.Name
    item.AltText       = body.AltText
    item.ImagePath     = body.ImagePath
    item.Type          = body.Type

    h.db.Save(&item)

    c.JSON(http.StatusOK, &item)
}

func (h *Handler) GetImage(c *gin.Context) {
    id := c.Param("id")
    var item models.Image
    if err := h.db.
        First(&item, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
        return
    }
    c.JSON(http.StatusOK, item)
}


type ImageListItem struct {
  Name          string `json:"name"`
  AltText       string `json:"alt_text"`
  ImagePath     string `json:"image_path"`
  Type          models.ImageType   `json:"type"`
//  ProductID     string `json:"product_id"`
}

func (h *Handler) ListImages(c *gin.Context) {
  var items []ImageListItem
  if err := h.db.Table("images").
    Select("name, alt_text, image_path, type").
    Scan(&items).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "images not found"})
  	return
  }
  c.JSON(http.StatusOK, gin.H{
  	"images": items,
  })
}

func (h *Handler) ListProductImages(c *gin.Context) {
  id := c.Param("id")  // Product_id in this case
  var items []ImageListItem
  if err := h.db.Model(&models.Image{}).
        Select("name, alt_text, image_path, type").
        Scan(&items).
        Where("product_id = ?", id).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Images not found"})
  	return
  }
  c.JSON(http.StatusOK, gin.H{
  	"images": items,
  })
}

func (h *Handler) DeleteImage(c *gin.Context) {
  id := c.Param("id")

  var item models.Image

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
