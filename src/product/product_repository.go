package product

import (
	"cashier-be/pkg/db"
	"cashier-be/pkg/models"
)

type IProductRepository interface {
	FindById(id uint) (models.Product, error)
	FindByIds(id []uint) ([]models.Product, error)
	FindByCode(code string) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id uint) error
	FindAll() ([]models.Product, error)
}

type ProductRepository struct {
	*db.IDbHandler
}

func (r *ProductRepository) FindById(id uint) (models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	return product, err
}

func (r *ProductRepository) FindByIds(ids []uint) ([]models.Product, error) {
	var products []models.Product
	if err := r.DB.Where("id IN ?", ids).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) FindByCode(code string) (models.Product, error) {
	var product models.Product
	err := r.DB.Where("code = ?", code).First(&product).Error
	return product, err
}

func (r *ProductRepository) Create(product models.Product) (models.Product, error) {
	err := r.DB.Create(&product).Error
	return product, err
}

func (r *ProductRepository) Update(product models.Product) (models.Product, error) {
	err := r.DB.Save(&product).Error
	return product, err
}

func (r *ProductRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Product{}, id).Error
	return err
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}
