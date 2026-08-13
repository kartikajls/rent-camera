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
	cameraHandler handler.CameraHandler,
	rentalOrderHandler handler.RentalOrderHandler,
	rentalOrderDetailHandler handler.RentalOrderDetailHandler,
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

	// Top Up endpoint
	user.POST("/topup", topUpHandler.Create)
	user.GET("/topup", topUpHandler.GetByUserID)
	user.GET("/topup/:id", topUpHandler.GetByID)

	// Camera endpoint
	user.GET("/cameras", cameraHandler.GetAll)
	user.GET("/cameras/:id", cameraHandler.GetByID)

	// Order endpoint
	user.POST("/orders", rentalOrderHandler.CreateOrder)
	user.GET("/orders", rentalOrderHandler.GetMyOrders)
	user.GET("/orders/:id", rentalOrderHandler.GetOrderByID)

	// Order Details endpoint
	user.POST("/rental-orders/:order_id/details", rentalOrderDetailHandler.CreateDetail)
	user.GET("/rental-orders/:order_id/details", rentalOrderDetailHandler.GetDetailsByOrderID)
	user.GET("/rental-order-details/:detail_id", rentalOrderDetailHandler.GetDetailByID)

	// =========================================
	// ADMIN ROUTES
	// =========================================

	admin := e.Group("/admin")

	admin.Use(
		middleware.JWTMiddleware,
		middleware.RoleMiddleware("admin"),
	)

	// Management
	admin.GET("/admin", userHandler.GetAll)

	// Top Up endpoint
	admin.GET("/topups", topUpHandler.GetAll)
	admin.GET("/topups/:id", topUpHandler.GetByID)
	admin.PUT("/topups/:id/approve", topUpHandler.Approve)
	admin.PUT("/topups/:id/reject", topUpHandler.Reject)

	// Camera endpoint
	admin.GET("/cameras", cameraHandler.GetAll)
	admin.GET("/cameras/:id", cameraHandler.GetByID)
	admin.POST("/cameras", cameraHandler.Create)
	admin.PUT("/cameras/:id", cameraHandler.Update)
	admin.DELETE("/cameras/:id", cameraHandler.Delete)

	// Order endpoint
	admin.GET("/orders", rentalOrderHandler.GetAllOrders)
	admin.PUT("/orders/:id/status", rentalOrderHandler.UpdateOrderStatus)
	admin.DELETE("/orders/:id", rentalOrderHandler.DeleteOrder)

	// Order Details endpoint
	admin.GET("/rental-order-details", rentalOrderDetailHandler.GetAllDetails)
	admin.PUT("/rental-order-details/:detail_id", rentalOrderDetailHandler.UpdateDetail)
	admin.DELETE("/rental-order-details/:detail_id", rentalOrderDetailHandler.DeleteDetail)

}
