package app

import (
	"doorProject/internal/config/configs"
	"doorProject/internal/db"
	v1 "doorProject/internal/delivery/http/v1"
	"doorProject/internal/delivery/http/v1/handlers"
	"doorProject/internal/delivery/http/v1/routes"
	"doorProject/internal/repository/psqlRepository"
	"doorProject/internal/server"
	"doorProject/internal/service"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	e := echo.New()
	v := validator.New(validator.WithRequiredStructEnabled())

	dbClient := db.NewDatabaseClient(dsn)
	configuredDB := configs.NewDatabaseConfig(dbClient.GetDBClient())
	configuredDB.SetupDb()

	productRepository := psqlRepository.NewProductRepository(configuredDB.Database)
	productService := service.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService, v)
	productRoutes := routes.NewProductRoute(productHandler)

	colorRepository := psqlRepository.NewColorRepository(configuredDB.Database)
	colorService := service.NewColorService(colorRepository)
	colorHandler := handlers.NewColorHandler(colorService, v)
	colorRoutes := routes.NewColorRoutes(colorHandler)

	categoryRepository := psqlRepository.NewCategoryRepository(configuredDB.Database)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryService, v)
	categoryRoutes := routes.NewCategoryRoute(categoryHandler)

	manyRoutes := v1.NewRoutes(productRoutes, colorRoutes, categoryRoutes, e)

	serv := server.NewServer(manyRoutes, e)
	serv.Start()
}
