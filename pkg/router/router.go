package router

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"github.com/slaghuis/YC-Product/pkg/handler"
  "github.com/slaghuis/YC-Product/pkg/config"
  "github.com/slaghuis/YC-Product/pkg/middleware"
)

func NewRouter(h *handler.Handler, cfg config.JwtConfigurations) *gin.Engine {
    r := gin.Default()

    // Health
    r.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })

    v1 := r.Group("/api/v1")

    // Example protected route for your own testing
    auth := v1.Group("/")
    auth.Use(middleware.RequireAccessToken(cfg))

    auth.GET("/test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": c.GetString("user_role")}) })

    auth.GET("/suppliers", h.ListSuppliers)
    auth.GET("/suppliers/:id", h.GetSupplier)
    auth.POST("/suppliers", h.CreateSupplier)
    auth.PUT("/suppliers/:id", h.UpdateSupplier)
    auth.DELETE("/suppliers/:id", h.DeleteSupplier)

    auth.GET("/categories", h.ListCategories)  // /products/categories
    auth.GET("/categories/parent", h.ListCategoriesByParent)   // /products/categories/parent?id=17 or &alt=lifestyle */
    auth.GET("/categories/:id", h.GetCategory)
    auth.POST("/categories", h.CreateCategory)
    auth.PUT("/categories/:id", h.UpdateCategory)
    auth.DELETE("/categories/:id", h.DeleteCategory)
    auth.GET("/categories/breadcrumb/:id", h.GetCategoryBreadcrumb)

    auth.GET("/:id/parameters", h.ListProductParameters)
    auth.PUT("/:id/parameters", h.UpdateProductParameters)

    auth.GET("/:id/categories", h.ListProductCategories)
    auth.PUT("/:id/categories", h.UpdateProductCategories)

    // Need to rethink this.  I want parameters for a product
    auth.GET("/images", h.ListImages)
    auth.GET("/:id/images", h.ListProductImages)
    auth.GET("/images/:id", h.GetImage)
    auth.POST("/images", h.CreateImage)
    auth.PUT("/images/:id", h.UpdateImage)
    auth.DELETE("/images/:id", h.DeleteImage)

    //auth.GET("/:id/like", h.GetLike)
    auth.PUT("/:id/like", h.ToggleLike)
    auth.GET("/:id/likes", h.CountLikes)

    auth.GET("/reviews", h.ListReviews)
    auth.GET("/:id/reviews", h.ListProductReviews)
    auth.GET("/reviews/:id", h.GetReview)
    auth.POST("/reviews", h.CreateReview).Use(middleware.RBACMiddleware(middleware.RoleCustomer, middleware.RoleAdmin, middleware.RoleSysAdmin))
    auth.PUT("/reviews/:id", h.UpdateReview).Use(middleware.RBACMiddleware(middleware.RoleCustomer, middleware.RoleAdmin, middleware.RoleSysAdmin))
    auth.DELETE("/reviews/:id", h.DeleteReview).Use(middleware.RBACMiddleware(middleware.RoleCustomer, middleware.RoleAdmin, middleware.RoleSysAdmin))

    auth.GET("/shop", h.ListProducts)
    auth.GET("/shop/count", h.CountProducts)
//    auth.GET("/shop_alt", h.ListShopProducts)
    auth.GET("/:id/shop", h.GetShopProduct)
    auth.GET("/:id/detail", h.GetShopProductDetail)

    auth.POST("/add", h.CreateProduct)

    auth.GET("/merchant", h.ListMerchantProducts)
    auth.GET("/:id/basic-info", h.GetProductBasicInfo)
    auth.PUT("/:id/basic-info", h.UpdateProductDefinition)

    auth.PUT("/merchant/:id/state", h.SetState)

    auth.GET("/merchant/suppliers", h.ListMerchantSuppliers)

    auth.GET("/merchant/categories", h.ListMerchantCategoriesByParent)

//    auth.PUT("/products/:id", h.UpdateProduct)  // NOT Needed.  We edit this bad boy in sections
//    auth.DELETE("/products/:id", h.DeleteProduct)  // TODO, but it will cause a ton of referential errors

    auth.POST("/:id/variants", h.CreateVariant)  // Create a variant and 0 or more options
    auth.GET("/:id/variants", h.GetVariantList)  //get product variants

    auth.DELETE("/variants/:id", h.DeleteVariant)
    auth.GET("/variants/:id", h.GetVariantOptions)  //get one variant and its options
    auth.POST("/variants/:id/options", h.CreateOption)

    auth.GET("/options/:id", h.GetOption)
    auth.PUT("/options/:id", h.UpdateOption)
    auth.DELETE("/options/:id", h.DeleteOption)

    auth.GET("/:id/skus", h.GetProductSKUHandler)
    auth.POST("/:id/skus", h.CreateSKUHandler)
    auth.POST("/:id/all_skus", h.CreateAllSKUsHandler)
    auth.GET("/:id/options_sku", h.GetSKUfromProductOptions)
//    auth.GET("/skus", h.ListSKUs)  //All SKU's - need pagination!
    auth.GET("/skus/:id", h.GetSKUHandler)  // Returns SKUListItem
    auth.PUT("/skus/:id", h.UpdateSKUHandler)  // Returns SKUListItem
    auth.DELETE("/skus/:id", h.DeleteSKUHandler)
    auth.GET("/skus/:id/product", h.GetProductFromSKU)

    return r
}
