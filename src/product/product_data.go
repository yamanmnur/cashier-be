package product

type ProductData struct {
	Id          uint
	Code        string
	Name        string
	Description string
	Barcode     string
	Price       float64
	Status      string
}

type ProductRequest struct {
	Code        string  `json:"code" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Barcode     string  `json:"barcode"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Status      string  `json:"status" validate:"required"`
}
