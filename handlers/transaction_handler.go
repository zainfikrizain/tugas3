package handlers

import (
	"Tugas3/models"
	"Tugas3/services"
	"net/http"

	"github.com/labstack/echo"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(s services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	var req models.CheckoutRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid  request",
		})
	}

	if len(req.Items) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "items cannot be empty",
		})
	}

	transaction, err := h.service.CreateTransaction(c.Request().context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, transaction)
}
