package transactions

import (
	"cashier-be/pkg/models"
	"cashier-be/src/product"
	"cashier-be/src/user"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ITransactionService interface {
	Create(userId uint, request *TransactionRequest) (TransactionData, error)
	List() ([]TransactionData, error)
	Detail(id uint) (TransactionData, error)
	Cancel(id uint) error
}

type TransactionService struct {
	TransactionRepository ITransactionRepository
	ProductRepo           product.IProductRepository
	UserRepo              user.IUserRepository
}

// --- Implementations ---

func (s *TransactionService) Create(userId uint, request *TransactionRequest) (TransactionData, error) {
	invoice := generateInvoiceNo()

	var total float64
	var items []models.TransactionItem

	var ids []uint
	for _, item := range request.Items {
		// convert string → uint
		ids = append(ids, item.ProductId)
	}

	products, err := s.ProductRepo.FindByIds(ids)
	if err != nil {
		return TransactionData{}, err
	}

	// build a map for faster lookup
	productMap := make(map[uint]models.Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	for _, reqItem := range request.Items {
		pid := reqItem.ProductId
		// get product from map
		product, ok := productMap[pid]
		if !ok {
			return TransactionData{}, fmt.Errorf("product not found: %d", pid)
		}

		itemTotal := product.Price * reqItem.Quantity
		total += itemTotal

		items = append(items, models.TransactionItem{
			Model:       gorm.Model{ID: 0},
			ProductId:   fmt.Sprintf("%d", product.ID), // keep as string if your model needs it
			Code:        product.Code,
			ProductName: product.Name,
			Barcode:     product.Barcode,
			Price:       product.Price,
			Quantity:    reqItem.Quantity,
			Total:       itemTotal,
		})
	}

	user, _ := s.UserRepo.FindById(userId)

	transaction := &models.Transaction{
		InvoiceNo:    invoice,
		UserId:       user.ID,
		UserName:     user.Username,
		CustomerName: request.CustomerName,
		PhoneNumber:  request.PhoneNumber,
		Total:        total, // convert ke string kalau di modelmu masih string
		Status:       models.PENDING,
		Items:        items,
	}

	created, err := s.TransactionRepository.Create(transaction)
	if err != nil {
		return TransactionData{}, err
	}

	return MapToTransactionData(*created), nil
}

func (s *TransactionService) List() ([]TransactionData, error) {
	transactions, err := s.TransactionRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var result []TransactionData
	for _, t := range transactions {
		result = append(result, MapToTransactionData(t))
	}
	return result, nil
}

func (s *TransactionService) Detail(id uint) (TransactionData, error) {
	transaction, err := s.TransactionRepository.FindByID(id)
	if err != nil {
		return TransactionData{}, err
	}
	return MapToTransactionData(*transaction), nil
}

func (s *TransactionService) Cancel(id uint) error {
	transaction, err := s.TransactionRepository.FindByID(id)
	if err != nil {
		return err
	}

	transaction.Status = models.CANCEL
	_, err = s.TransactionRepository.Update(transaction)
	return err
}

// --- Helpers ---

func generateInvoiceNo() string {
	today := time.Now().Format("20060102") // yyyymmdd
	// NOTE: Sementara pakai timestamp, bisa diganti auto increment kalau perlu
	return fmt.Sprintf("INV-%s-%d", today, time.Now().Unix())
}
