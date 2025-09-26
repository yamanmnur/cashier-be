package routes

import (
	"cashier-be/pkg/db"
	"cashier-be/src/reports"
	"cashier-be/src/transactions"

	"github.com/labstack/echo/v4"
)

func ReportRoute(dbHandler *db.IDbHandler, e *echo.Group) {
	// Init repository
	transactionRepo := transactions.TransactionRepository{IDbHandler: dbHandler}

	// Init service
	reportService := reports.ReportService{
		TransactionRepo: &transactionRepo,
	}

	// Init handler
	reportHandler := reports.ReportHandler{Service: &reportService}

	// Define route group
	reportRoute := e.Group("/report")

	// Routes
	reportRoute.GET("/pdf", reportHandler.ExportPDF)
	reportRoute.GET("/excel", reportHandler.ExportExcel)
}
