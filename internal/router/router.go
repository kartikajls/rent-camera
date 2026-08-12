package router

import (
	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/internal/handler"
	"p2-ip-kartikajls/internal/middleware"
)

func SetupRoutes(
	e *echo.Echo,
	userHandler handler.UserHandler,
	topUpHandler handler.TopUpHandler,
) {

	// =========================================
	// PUBLIC ROUTES
	// =========================================

	e.POST("/users/register", userHandler.Register)
	e.POST("/users/login", userHandler.Login)

	// =========================================
	// USER ROUTES
	// =========================================

	user := e.Group("/users")

	user.Use(
		middleware.JWTMiddleware,
		middleware.RoleMiddleware("user"),
	)

	user.GET("/:id", userHandler.GetByID)

	// Top Up User
	user.POST("/topup", topUpHandler.Create)
	user.GET("/topup", topUpHandler.GetByUserID)
	user.GET("/topup/:id", topUpHandler.GetByID)

	// =========================================
	// ADMIN ROUTES
	// =========================================

	admin := e.Group("/admin")

	admin.Use(
		middleware.JWTMiddleware,
		middleware.RoleMiddleware("admin"),
	)

	// Management
	admin.GET("/users", userHandler.GetAll)

	// Top Up Management
	admin.GET("/topups", topUpHandler.GetAll)
	admin.GET("/topups/:id", topUpHandler.GetByID)
	admin.PUT("/topups/:id/approve", topUpHandler.Approve)
	admin.PUT("/topups/:id/reject", topUpHandler.Reject)
}
