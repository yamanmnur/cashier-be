package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		secretKey := viper.GetString("APP_SECRET_KEY")
		if secretKey == "" {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "secret key not configured",
			})
		}
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Missing or invalid token",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse JWT
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "Invalid token",
			})
		}

		// Extract subject (userId) if present
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, ok := claims["sub"].(string); ok {
				c.Set("UserId", sub)
				fmt.Println("FROM MIDDLEWARE", sub)
			}
		}

		return next(c)
	}
}
