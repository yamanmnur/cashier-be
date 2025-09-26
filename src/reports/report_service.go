package reports

import (
	"bytes"
	"cashier-be/src/transactions"
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
)

type IReportService interface {
	ExportPDF() ([]byte, error)
	ExportExcel() ([]byte, error)
}

type ReportService struct {
	TransactionRepo transactions.ITransactionRepository
}

// ExportPDF generates a PDF report of all transactions
func (s *ReportService) ExportPDF() ([]byte, error) {
	transactions, err := s.TransactionRepo.FindAll()
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Transaction Report")

	pdf.Ln(12)
	pdf.SetFont("Arial", "", 10)

	for _, tx := range transactions {
		pdf.CellFormat(0, 8, fmt.Sprintf("Invoice: %s | Customer: %s | Total: %.2f | Status: %s",
			tx.InvoiceNo, tx.CustomerName, tx.Total, tx.Status), "", 1, "", false, 0, "")
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ExportExcel generates an Excel report of all transactions
func (s *ReportService) ExportExcel() ([]byte, error) {
	transactions, err := s.TransactionRepo.FindAll()
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Report"
	f.NewSheet(sheet)

	// Header
	headers := []string{"InvoiceNo", "CustomerName", "PhoneNumber", "Total", "Status"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Data
	for rowIdx, tx := range transactions {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", rowIdx+2), tx.InvoiceNo)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", rowIdx+2), tx.CustomerName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", rowIdx+2), tx.PhoneNumber)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", rowIdx+2), tx.Total)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", rowIdx+2), tx.Status)
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
