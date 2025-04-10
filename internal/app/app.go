package app

import (
	"context"
	"doorProject/internal/config/configs"
	"doorProject/internal/db"
	"doorProject/internal/delivery/http/v1"
	"doorProject/internal/delivery/http/v1/handlers"
	"doorProject/internal/delivery/http/v1/routes"
	"doorProject/internal/repository/psqlRepository"
	"doorProject/internal/server"
	"doorProject/internal/service"
	"doorProject/pkg/config"
	messageHandlerService "doorProject/pkg/service"
	"doorProject/pkg/service/QueueHandlers"
	"doorProject/pkg/service/QueueManager"
	"doorProject/pkg/service/jwtAuth"
	"doorProject/pkg/service/smtpSender"
	"log"
	"os"
	"strconv"

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

	jwtSecret := os.Getenv("JWT_SECRET")
	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	accessCookie := os.Getenv("ACCESS_TOKEN_COOKIE_NAME")
	refreshCookie := os.Getenv("REFRESH_TOKEN_COOKIE_NAME")
	tokenExpiration, err := strconv.ParseInt(os.Getenv("TOKEN_EXPIRATION_TIME"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	refreshExpiration, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXPIRATION_TIME"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	redisDsn := os.Getenv("REDIS_URL")
	mailhogClient := os.Getenv("MAILER_DSN")
	smtpFrom := os.Getenv("MAILER_FROM")

	dbClient := db.NewDatabaseClient(dsn)
	configuredDB := configs.NewDatabaseConfig(dbClient.GetDBClient())
	configuredDB.SetupDb()

	tokenRepo := psqlRepository.NewRefreshTokenRepository(configuredDB.Database)
	workerRepository := psqlRepository.NewWorkerRepository(configuredDB.Database)

	authService := jwtAuth.NewJWTAuthService(
		accessCookie,
		refreshCookie,
		jwtSecret,
		refreshSecret,
		int(tokenExpiration),
		int(refreshExpiration),
		tokenRepo,
		workerRepository,
	)

	jwtConfig := &configs.MiddlewareConfig{
		SigningKey:   []byte(jwtSecret),
		TokenLookup:  "cookie:access-token",
		ErrorHandler: authService.JWTErrorChecker,
	}

	redisClient := config.NewRedisClient(redisDsn)
	smtpClient := smtpSender.NewMailhogClient(mailhogClient, smtpFrom)
	mailSender := smtpSender.NewSenderService(smtpClient)

	subHandlers := make(map[string]messageHandlerService.MessageHandlerInterface)
	subscriber := QueueManger.NewRedisSubscriber(redisClient, subHandlers)

	emailSenderHandler := QueueHandlers.NewEmailSenderHandler(redisClient, mailSender)

	subscriber.RegisterHandler("sendEmail", emailSenderHandler)

	ctx := context.Background()
	topics := []string{"sendEmail"}
	subscriber.ConsumeMessages(ctx, topics)

	publisher := QueueManger.NewRedisPublisher(redisClient)

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

	workerService := *service.NewWorkerService(workerRepository)
	workerHandlers := handlers.NewWorkerHandlers(workerService, *v)
	workerRoutes := routes.NewWorkerRoutes(*workerHandlers)

	workerCalendarRepository := psqlRepository.NewWorkerCalendarRepository(configuredDB.Database)
	workerCalendarService := service.NewWorkerCalendar(workerCalendarRepository)
	workerCalendarHandlers := handlers.NewWorkerCalendarHandlers(workerCalendarService, v)
	workerCalendarRoutes := routes.NewWorkerCalendar(workerCalendarHandlers, jwtConfig, authService)

	clientRepository := psqlRepository.NewClientRepository(configuredDB.Database)
	clientService := service.NewClientService(clientRepository)
	clientHandlers := handlers.NewClientHandlers(clientService, v)
	clientRoutes := routes.NewClientRoutes(clientHandlers)

	serviceRepository := psqlRepository.NewServiceRepository(configuredDB.Database)
	serviceService := service.NewServiceService(serviceRepository)
	serviceHandler := handlers.NewServiceHandler(serviceService, v)
	serviceRoutes := routes.NewServiceRoutes(serviceHandler)

	cartRepository := psqlRepository.NewCartRepository(configuredDB.Database)
	cartService := service.NewCartService(cartRepository)
	cartHandlers := handlers.NewCartHandlers(cartService, v)
	cartRoutes := routes.NewCartRoutes(cartHandlers)

	appointmentRepository := psqlRepository.NewAppointmentRepository(configuredDB.Database)
	appointmentService := service.NewAppointmentService(
		workerCalendarRepository,
		appointmentRepository,
		serviceRepository,
	)
	appointmentHandlers := handlers.NewAppointmentHandler(appointmentService, v)
	appointmentRoutes := routes.NewAppointmentRoutes(appointmentHandlers)

	authHandlers := handlers.NewAuthHandlers(authService, workerService, publisher)
	authRoutes := routes.NewAuthRoutes(authHandlers)

	manyRoutes := v1.NewRoutes(
		productRoutes,
		colorRoutes,
		categoryRoutes,
		workerRoutes,
		workerCalendarRoutes,
		clientRoutes,
		serviceRoutes,
		cartRoutes,
		appointmentRoutes,
		authRoutes,
		e,
	)

	serv := server.NewServer(manyRoutes, e)
	serv.Start()
}
