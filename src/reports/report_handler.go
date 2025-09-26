package reports

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReportHandler struct {
	Service IReportService
}

func (h *ReportHandler) ExportPDF(c echo.Context) error {
	data, err := h.Service.ExportPDF()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.Blob(http.StatusOK, "application/pdf", data)
}

func (h *ReportHandler) ExportExcel(c echo.Context) error {
	data, err := h.Service.ExportExcel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.Blob(http.StatusOK,
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		data,
	)
}
