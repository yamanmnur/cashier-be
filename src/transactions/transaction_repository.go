package transactions

import (
	"cashier-be/pkg/db"
	"cashier-be/pkg/models"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Create(transaction *models.Transaction) (*models.Transaction, error)
	FindByID(id uint) (*models.Transaction, error)
	FindByInvoice(invoiceNo string) (*models.Transaction, error)
	FindAll() ([]models.Transaction, error)
	Update(transaction *models.Transaction) (*models.Transaction, error)
	Delete(id uint) error
}

type TransactionRepository struct {
	*db.IDbHandler
}

func (r *TransactionRepository) Create(transaction *models.Transaction) (*models.Transaction, error) {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.DB.Preload("Items").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindByInvoice(invoiceNo string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.DB.Preload("Items").Where("invoice_no = ?", invoiceNo).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Preload("Items").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) Update(transaction *models.Transaction) (*models.Transaction, error) {
	err := r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) Delete(id uint) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// Delete items first (FK)
		if err := tx.Where("transaction_id = ?", id).Delete(&models.TransactionItem{}).Error; err != nil {
			return err
		}
		// Delete parent
		if err := tx.Delete(&models.Transaction{}, id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
