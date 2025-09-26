package main

import (
	"cashier-be/pkg/db"
	"cashier-be/routes"
	"flag"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	flag.Parse()

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	dbUrl := viper.Get("DB_URL").(string)
	APP_PORT := viper.Get("APP_PORT").(string)

	dbPgsql, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	dbHandler := db.Init(dbPgsql)
	db.InitInstanceDbHandler(&db.IDbHandler{
		DB: dbHandler,
	})

	e := echo.New()
	// Custom CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},                                      // Specify allowed origins
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},                   // Specify allowed methods
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}, // Specify allowed headers
		// AllowCredentials: true, // Set to true if you need to send cookies with cross-origin requests
	}))

	api := e.Group("/api/v1")

	routes.AuthRoute(db.GetDbHandler(), api)
	routes.ProductRoute(db.GetDbHandler(), api)
	routes.TransactionRoute(db.GetDbHandler(), api)
	routes.ReportRoute(db.GetDbHandler(), api)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Cashier API")
	})

	e.Logger.Fatal(e.Start("127.0.0.1:" + APP_PORT))
}
