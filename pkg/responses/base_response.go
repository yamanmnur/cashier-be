package pkg_response

import pkg_data "cashier-be/pkg/data"

type MetaData struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type GenericResponse struct {
	MetaData MetaData    `json:"meta_data"`
	Data     interface{} `json:"data"`
}

type BasicResponse struct {
	MetaData MetaData `json:"meta_data"`
}

type PaginateResponse[T any] struct {
	MetaData MetaData          `json:"meta_data"`
	Data     []T               `json:"data"`
	PageData pkg_data.PageData `json:"page_data"`
}

func NewPaginateResponse[T any](metaData MetaData) *PaginateResponse[T] {
	return &PaginateResponse[T]{
		MetaData: metaData,
	}
}

func (p *PaginateResponse[T]) GetPageData() pkg_data.PageData {
	return p.PageData
}

func (p *PaginateResponse[T]) SetPageData(pageData pkg_data.PageData) {
	p.PageData = pageData
}

func (p *PaginateResponse[T]) GetData() []T {
	return p.Data
}

func (p *PaginateResponse[T]) SetData(data []T) {
	p.Data = data
}

// Success response with data
func SuccessResponse(message string, data interface{}) GenericResponse {
	return GenericResponse{
		MetaData: MetaData{
			Message: message,
			Code:    "200",
		},
		Data: data,
	}
}

func ErrorResponse(message string, err error) GenericResponse {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return GenericResponse{
		MetaData: MetaData{
			Message: message,
			Code:    "400", // you can override per handler if needed
		},
		Data: map[string]string{
			"error": errMsg,
		},
	}
}
