package controller

import (
	"net/http"
	"MiniProject/model"
	"strconv"

	"MiniProject/usecase"

	"github.com/labstack/echo/v4"
)

func GetTransactionscontroller(c echo.Context) error {
	transactions, e := usecase.GetListTransactions()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"transactions":  transactions,
	})
}

func GetTransactionController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	transaction, err := usecase.GetTransaction(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get transaction",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"transaction":  transaction,
	})
}

// create new book
func CreateTransactionController(c echo.Context) error {
	transaction := model.Transaction{}
	c.Bind(&transaction)

	if err := usecase.CreateTransaction(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create transaction",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new transaction",
		"transaction":    transaction,
	})
}

// delete book by id
func DeleteTransactionController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteTransaction(uint(id)); err != nil {
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

// update book by id
func UpdateTransactionController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	transaction := model.Transaction{}
	c.Bind(&transaction)
	transaction.ID = uint(id)
	if err := usecase.UpdateTransaction(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update transaction",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf transaction tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update transaction",
	})
}
