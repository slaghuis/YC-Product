package main

import (
  "context"
  "log"
  "os"
  "time"
  "github.com/joho/godotenv"
  "github.com/slaghuis/YC-Product/pkg/api"
  "github.com/slaghuis/YC-Product/pkg/config"
  "github.com/slaghuis/YC-Product/pkg/clients"
  "github.com/slaghuis/YC-Product/pkg/database"
  "github.com/slaghuis/YC-Product/pkg/gateway"
  "github.com/slaghuis/YC-Product/pkg/handler"
  "github.com/slaghuis/YC-Product/pkg/router"
  "github.com/slaghuis/YC-Product/pkg/repositories"
  "github.com/slaghuis/YC-Product/pkg/services"

)

var cfg config.Configurations

func init() {
  // Set the file name of the configurations file
  mode := os.Getenv("GIN_MODE")
  if len(mode) == 0 {
    log.Println("GIN_MODE env var not set, assuming development")
    mode = "development"
  }
  log.Printf("Starting Micro Service in %s mode", mode)
}

func main() {
  // Load environment variables from .env file
  err := godotenv.Load()
  if err != nil {
    log.Fatalf("Error loading .env file: %v", err)
  }
  cfg = config.LoadConfig()

  // Define a API object
  api :=  api.NewAPIServer("http://localhost:8080") // URL of the API gateway

  // Initialize Database
  db, err := database.Connect(cfg.Database)
  if err != nil {
    log.Fatal(err)
    panic("Cannot connect to DB")
  }
  log.Println("Connected to Database!")

  // ProductRepository → products, variants, options, categories
  productRepo := repositories.NewProductRepository(db)
  skuRepo := repositories.NewSkuRepository(db)
  supplierRepo := repositories.NewSupplierRepository(db)

  // Move this to another microservice!
  reviewRepo := repositories.NewReviewRepository(db)

  priceCacheRepo := repositories.NewPriceCacheRepository(db)
  imageCacheRepo := repositories.NewImageCacheRepository(db)

  searchRepo := repositories.NewSearchRepository()

  identityClient := clients.NewIdentityClient(api)
  pricingClient := clients.NewPricingClient(api)
  preferenceClient := clients.NewPreferenceClient(api)
  imagesClient := clients.NewImagesClient(api)

  skuSvc := services.NewSkuService(skuRepo, priceCacheRepo, productRepo, identityClient, pricingClient)
  searchSvc := services.NewProductSearchService(productRepo, searchRepo)
  productSvc := services.NewProductService(productRepo, imageCacheRepo, searchSvc, skuSvc, preferenceClient, imagesClient)  //Inject the whole service as a dependency!
  supplierSvc :=services.NewSupplierService(supplierRepo, preferenceClient)

  reviewSvc :=services.NewReviewService(reviewRepo)   // move this to another Micro-service

  handler := handler.NewHandler(db, api, skuSvc, productSvc, supplierSvc, reviewSvc)
  handler.Migrate()

  // Index the products to allow faceted search
  go func() {
    var ctx context.Context
    err := searchSvc.IndexProducts(ctx)
    if err != nil {
      log.Fatal(err)
      panic("Cannot initialise search index")
    }
  } ()

  // Register after a slight delay to ensure server is up
  go func() {
      time.Sleep(2 * time.Second)
      gateway.RegisterWithGateway()
  }()

  router := router.NewRouter(handler, cfg.Security)
  router.Run(":" + cfg.Server.Port)
}
