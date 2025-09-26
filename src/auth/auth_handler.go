package auth

import (
	pkg_response "cashier-be/pkg/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service IAuthService
}

func (h *AuthHandler) Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	res, err := h.Service.Login(req)

	if err != nil {
		response := pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: err.Error(),
				Code:    "200",
			},
		}

		return c.JSON(http.StatusUnauthorized, response)

	}

	response := pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Message: "Success To Login User",
			Code:    "200",
		},
		Data: res,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) Register(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	res, err := h.Service.Register(req)

	if err != nil {
		response := pkg_response.GenericResponse{
			MetaData: pkg_response.MetaData{
				Message: err.Error(),
				Code:    "200",
			},
			Data: res,
		}

		return c.JSON(http.StatusInternalServerError, response)
	}

	response := pkg_response.GenericResponse{
		MetaData: pkg_response.MetaData{
			Message: "Success To Register User",
			Code:    "200",
		},
		Data: res,
	}

	return c.JSON(http.StatusOK, response)
}
