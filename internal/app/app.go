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

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	e := echo.New()

	dbClient := db.NewDatabaseClient(dsn)
	configuredDB := configs.NewDatabaseConfig(dbClient.GetDBClient())
	productRepository := psqlRepository.NewProductRepository(configuredDB.Database)
	colorRepository := psqlRepository.NewColorRepository(configuredDB.Database)

	productService := service.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)
	productRoutes := routes.NewProductRoute(productHandler)

	colorService := service.NewColorService(colorRepository)
	colorHandler := handlers.NewColorHandler(colorService)
	colorRoutes := routes.NewColorRoutes(colorHandler)

	manyRoutes := v1.NewRoutes(productRoutes, colorRoutes, e)

	serv := server.NewServer(manyRoutes, e)
	serv.Start()
}
