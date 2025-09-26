package transactions

import (
	"cashier-be/pkg/models"
)

// --- Request DTO ---
type TransactionItemRequest struct {
	ProductId uint    `json:"product_id" validate:"required"`
	Quantity  float64 `json:"quantity" validate:"required,gt=0"`
}

type TransactionRequest struct {
	CustomerName string                   `json:"customer_name" validate:"required"`
	PhoneNumber  string                   `json:"phone_number"`
	Items        []TransactionItemRequest `json:"items" validate:"required,dive"`
}

// --- Response DTO ---
type TransactionItemData struct {
	ID          uint    `json:"id"`
	ProductId   string  `json:"product_id"`
	Code        string  `json:"code"`
	ProductName string  `json:"product_name"`
	Barcode     string  `json:"barcode"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
	Total       float64 `json:"total"`
}

type TransactionData struct {
	ID           uint                  `json:"id"`
	InvoiceNo    string                `json:"invoice_no"`
	UserName     string                `json:"user_name"`
	CustomerName string                `json:"customer_name"`
	PhoneNumber  string                `json:"phone_number"`
	Total        float64               `json:"total"`
	Status       models.Status         `json:"status"`
	Items        []TransactionItemData `json:"items"`
}

func MapToTransactionData(t models.Transaction) TransactionData {
	var items []TransactionItemData
	for _, i := range t.Items {
		items = append(items, TransactionItemData{
			ID:          i.ID,
			ProductId:   i.ProductId,
			Code:        i.Code,
			ProductName: i.ProductName,
			Barcode:     i.Barcode,
			Price:       i.Price,
			Quantity:    i.Quantity,
			Total:       i.Total,
		})
	}

	return TransactionData{
		ID:           t.ID,
		InvoiceNo:    t.InvoiceNo,
		UserName:     t.UserName,
		CustomerName: t.CustomerName,
		PhoneNumber:  t.PhoneNumber,
		Total:        t.Total,
		Status:       t.Status,
		Items:        items,
	}
}
