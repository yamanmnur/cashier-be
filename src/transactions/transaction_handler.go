package transactions

import (
	pkg_response "cashier-be/pkg/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	Service ITransactionService
}

// --- Create Transaction ---
func (h *TransactionHandler) Create(c echo.Context) error {
	var req TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, pkg_response.ErrorResponse("Invalid request payload", err))
	}

	// get userId from JWT middleware (set via c.Set / c.Locals)
	stringId := fmt.Sprintf("%v", c.Get("UserId"))

	userId, err := strconv.ParseUint(stringId, 10, 32)
	fmt.Println("USER ID", userId)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, pkg_response.ErrorResponse("Invalid user ID type", nil))
	}

	tx, err := h.Service.Create(uint(userId), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_response.ErrorResponse("Failed to create transaction", err))
	}

	return c.JSON(http.StatusCreated, pkg_response.SuccessResponse("Transaction created", tx))
}

// --- List Transactions ---
func (h *TransactionHandler) List(c echo.Context) error {
	txs, err := h.Service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_response.ErrorResponse("Failed to fetch transactions", err))
	}
	return c.JSON(http.StatusOK, pkg_response.SuccessResponse("Transaction list", txs))
}

// --- Detail Transaction ---
func (h *TransactionHandler) Detail(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg_response.ErrorResponse("Invalid transaction ID", err))
	}

	tx, err := h.Service.Detail(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, pkg_response.ErrorResponse("Transaction not found", err))
	}

	return c.JSON(http.StatusOK, pkg_response.SuccessResponse("Transaction detail", tx))
}

// --- Cancel Transaction ---
func (h *TransactionHandler) Cancel(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg_response.ErrorResponse("Invalid transaction ID", err))
	}

	if err := h.Service.Cancel(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, pkg_response.ErrorResponse("Failed to cancel transaction", err))
	}

	return c.JSON(http.StatusOK, pkg_response.SuccessResponse("Transaction canceled", nil))
}
