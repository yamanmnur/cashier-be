package product

import "cashier-be/pkg/models"

type IProductService interface {
	Create(request *ProductRequest) (ProductData, error)
	Update(productId uint, request *ProductRequest) (ProductData, error)
	List() ([]ProductData, error)
	Delete(productId uint) error
	Detail(productId uint) (ProductData, error)
}

type ProductService struct {
	ProductRepository IProductRepository
}

func NewProductService(repo IProductRepository) IProductService {
	return &ProductService{
		ProductRepository: repo,
	}
}

func (s *ProductService) Create(request *ProductRequest) (ProductData, error) {
	product := models.Product{
		Code:        request.Code,
		Barcode:     request.Barcode,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Status:      request.Status,
	}

	created, err := s.ProductRepository.Create(product)
	if err != nil {
		return ProductData{}, err
	}

	return mapToProductData(created), nil
}

func (s *ProductService) Update(productId uint, request *ProductRequest) (ProductData, error) {
	// Find existing product
	existing, err := s.ProductRepository.FindById(productId)
	if err != nil {
		return ProductData{}, err
	}

	// Update fields
	existing.Code = request.Code
	existing.Name = request.Name
	existing.Price = request.Price
	existing.Barcode = request.Barcode
	existing.Code = request.Code
	existing.Status = request.Status
	existing.Description = request.Description

	updated, err := s.ProductRepository.Update(existing)
	if err != nil {
		return ProductData{}, err
	}

	return mapToProductData(updated), nil
}

func (s *ProductService) List() ([]ProductData, error) {
	var result []ProductData
	// You may need to implement FindAll in repository
	products, err := s.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		result = append(result, mapToProductData(p))
	}

	return result, nil
}

func (s *ProductService) Delete(productId uint) error {
	return s.ProductRepository.Delete(productId)
}

func (s *ProductService) Detail(productId uint) (ProductData, error) {
	product, err := s.ProductRepository.FindById(productId)
	if err != nil {
		return ProductData{}, err
	}
	return mapToProductData(product), nil
}

// --- Mapper helper ---
func mapToProductData(product models.Product) ProductData {
	return ProductData{
		Id:          product.ID,
		Code:        product.Code,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      product.Status,
		Barcode:     product.Barcode,
	}
}
