package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/config"
	"p2-ip-kartikajls/internal/handler"
	"p2-ip-kartikajls/internal/repository"
	"p2-ip-kartikajls/internal/router"
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

	// Repository
	userRepository := repository.NewUserRepository(db)
	topUpRepository := repository.NewTopUpRepository(db)
	cameraRepository := repository.NewCameraRepository(db)
	rentOrderRepository := repository.NewRentalOrderRepository(db)

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepository)
	topUpUsecase := usecase.NewTopUpUsecase(topUpRepository, userRepository)
	cameraUsecase := usecase.NewCameraUsecase(cameraRepository)
	rentOrderUsecase := usecase.NewRentalOrderUsecase(rentOrderRepository)

	// Handler
	userHandler := handler.NewUserHandler(userUsecase)
	topUpHandler := handler.NewTopUpHandler(topUpUsecase)
	cameraHandler := handler.NewCameraHandler(cameraUsecase)
	rentOrderHandler := handler.NewRentalOrderHandler(rentOrderUsecase)

	// Echo
	e := echo.New()

	// Router
	router.SetupRoutes(
		e,
		userHandler,
		topUpHandler,
		cameraHandler,
		rentOrderHandler,
	)

	// Start Server
	log.Println("Server running on :8080")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
