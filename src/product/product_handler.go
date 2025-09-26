package product

import (
	pkg_response "cashier-be/pkg/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Service IProductService
}

func (h *ProductHandler) List(c echo.Context) error {
	products, err := h.Service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: err.Error(),
				Code:    "500",
			},
		})
	}

	return c.JSON(http.StatusOK, pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Message: "Success get product list",
			Code:    "200",
		},
		Data: products,
	})
}

func (h *ProductHandler) Create(c echo.Context) error {
	var req ProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: "Invalid request body",
				Code:    "400",
			},
		})
	}

	product, err := h.Service.Create(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: err.Error(),
				Code:    "500",
			},
		})
	}

	return c.JSON(http.StatusCreated, pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Message: "Product created successfully",
			Code:    "201",
		},
		Data: product,
	})
}

func (h *ProductHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: "Invalid product ID",
				Code:    "400",
			},
		})
	}

	var req ProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: "Invalid request body",
				Code:    "400",
			},
		})
	}

	product, err := h.Service.Update(uint(id), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: err.Error(),
				Code:    "500",
			},
		})
	}

	return c.JSON(http.StatusOK, pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Message: "Product updated successfully",
			Code:    "200",
		},
		Data: product,
	})
}

func (h *ProductHandler) Detail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: "Invalid product ID",
				Code:    "400",
			},
		})
	}

	product, err := h.Service.Detail(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: err.Error(),
				Code:    "500",
			},
		})
	}

	return c.JSON(http.StatusOK, pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Message: "Success get product detail",
			Code:    "200",
		},
		Data: product,
	})
}

func (h *ProductHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: "Invalid product ID",
				Code:    "400",
			},
		})
	}

	if err := h.Service.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: err.Error(),
				Code:    "500",
			},
		})
	}

	return c.JSON(http.StatusOK, pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Message: "Product deleted successfully",
			Code:    "200",
		},
	})
}
