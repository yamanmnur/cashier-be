package routes

import (
	"cashier-be/pkg/db"
	"cashier-be/src/auth"
	"cashier-be/src/user"

	"github.com/labstack/echo/v4"
)

func AuthRoute(dbHandler *db.IDbHandler, e *echo.Group) {
	repository := user.UserRepository{IDbHandler: dbHandler}
	authService := auth.AuthService{UserRepository: &repository}
	authHandler := auth.AuthHandler{Service: &authService}

	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)
}
