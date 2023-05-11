package controller

import (
	"miniproject/model"
	"miniproject/repository/database"
	"miniproject/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController interface {
	GetTransactioncontroller(c echo.Context) error
	GetTransactionController(c echo.Context) error
	CreateTransactionController(c echo.Context) error
	DeleteTransactionController(c echo.Context) error
	UpdateTransactionController(c echo.Context) error
}

type transactionController struct {
	transactionUsecase    usecase.TransactionUsecase
	transactionRepository database.TransactionRepository
}

func NewTransactionController(
	transactionUsecase usecase.TransactionUsecase,
	transactionRepository database.TransactionRepository,
) *transactionController {
	return &transactionController{
		transactionUsecase,
		transactionRepository,
	}
}

func (b *transactionController) GetTransactioncontroller(c echo.Context) error {
	transactions, e := b.transactionRepository.GetTransactions()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":       "success",
		"transactions": transactions,
	})
}

func (b *transactionController) GetTransactionController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	transaction, err := b.transactionRepository.GetTransaction(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get transaction",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":       "success",
		"transactions": transaction,
	})
}

// create new transaction
func (b *transactionController) CreateTransactionController(c echo.Context) error {
	transaction := model.Transaction{}
	c.Bind(&transaction)

	if err := b.transactionUsecase.CreateTransaction(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create transaction",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create new transaction",
		"transaction": transaction,
	})
}

// delete transaction by id
func (b *transactionController) DeleteTransactionController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := b.transactionUsecase.DeleteTransaction(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete transaction",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf transaction tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete transaction",
	})
}

// update transaction by id
func (b *transactionController) UpdateTransactionController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	transaction := model.Transaction{}
	c.Bind(&transaction)
	transaction.ID = uint(id)
	if err := b.transactionUsecase.UpdateTransaction(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update transaction",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf alat tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update transaction",
	})
}
