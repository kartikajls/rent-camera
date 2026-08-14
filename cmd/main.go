package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/config"
	"p2-ip-kartikajls/internal/handler"
	"p2-ip-kartikajls/internal/repository"
	"p2-ip-kartikajls/internal/router"
	"p2-ip-kartikajls/internal/service"
	"p2-ip-kartikajls/internal/usecase"
)

func main() {

	// Load Environment
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	// Database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// WA SERVICE
	whatsappService := service.NewWhatsAppService()

	// Repository
	userRepository := repository.NewUserRepository(db)
	topUpRepository := repository.NewTopUpRepository(db)
	cameraRepository := repository.NewCameraRepository(db)
	orderRepository := repository.NewRentalOrderRepository(db)
	orderDetailsRepository := repository.NewRentalOrderDetailRepository(db)
	paymentRepository := repository.NewPaymentRepository(db)

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepository)
	topUpUsecase := usecase.NewTopUpUsecase(topUpRepository, userRepository, whatsappService)
	cameraUsecase := usecase.NewCameraUsecase(cameraRepository)
	orderUsecase := usecase.NewRentalOrderUsecase(orderRepository, cameraRepository, userRepository, whatsappService)
	orderDetailsUsecase := usecase.NewRentalOrderDetailUsecase(orderDetailsRepository)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepository, userRepository, orderRepository, whatsappService)
	// Usecase khusus testing WhatsApp
	wasenderUsecase := usecase.NewWasenderUsecase()

	// Handler
	userHandler := handler.NewUserHandler(userUsecase)
	topUpHandler := handler.NewTopUpHandler(topUpUsecase)
	cameraHandler := handler.NewCameraHandler(cameraUsecase)
	orderHandler := handler.NewRentalOrderHandler(orderUsecase)
	orderDetailsHandler := handler.NewRentalOrderDetailHandler(orderDetailsUsecase)
	paymentHandler := handler.NewPaymentHandler(paymentUsecase)
	wasenderHandler := handler.NewWasenderHandler(wasenderUsecase)

	// Echo
	e := echo.New()

	// Router
	router.SetupRoutes(
		e,
		userHandler,
		topUpHandler,
		cameraHandler,
		orderHandler,
		orderDetailsHandler,
		paymentHandler,
		wasenderHandler,
	)

	// Start Server
	log.Println("Server running on :8080")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
