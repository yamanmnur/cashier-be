package routes

import (
	"cashier-be/pkg/db"
	"cashier-be/pkg/middlewares"
	"cashier-be/src/product"

	"github.com/labstack/echo/v4"
)

func ProductRoute(dbHandler *db.IDbHandler, e *echo.Group) {
	repository := product.ProductRepository{IDbHandler: dbHandler}
	productService := product.ProductService{ProductRepository: &repository}
	productHandler := product.ProductHandler{Service: &productService}

	productRoute := e.Group("/master/product", middlewares.JwtMiddleware)

	// Routes
	productRoute.GET("/list", productHandler.List)
	productRoute.GET("/:id", productHandler.Detail)
	productRoute.POST("", productHandler.Create)
	productRoute.PUT("/:id", productHandler.Update)
	productRoute.DELETE("/:id", productHandler.Delete)
}
