package routes

import (
	"cashier-be/pkg/db"
	"cashier-be/pkg/middlewares"
	"cashier-be/src/product"
	"cashier-be/src/transactions"
	"cashier-be/src/user"

	"github.com/labstack/echo/v4"
)

func TransactionRoute(dbHandler *db.IDbHandler, e *echo.Group) {
	repository := transactions.TransactionRepository{IDbHandler: dbHandler}
	productRepository := product.ProductRepository{IDbHandler: dbHandler}
	userRepository := user.UserRepository{IDbHandler: dbHandler}
	service := transactions.TransactionService{TransactionRepository: &repository, ProductRepo: &productRepository, UserRepo: &userRepository}
	handler := transactions.TransactionHandler{Service: &service}

	// Define route group
	txRoute := e.Group("/transactions", middlewares.JwtMiddleware)

	// Routes
	txRoute.POST("", handler.Create)           // Create transaction
	txRoute.GET("", handler.List)              // List all transactions
	txRoute.GET("/:id", handler.Detail)        // Get transaction detail
	txRoute.PUT("/:id/cancel", handler.Cancel) // Cancel transaction
}
