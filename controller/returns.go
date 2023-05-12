package controller

import (
	"net/http"
	"MiniProject/model"
	"MiniProject/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetReturnssController(c echo.Context) error {
	returnss, e := usecase.GetListReturnss()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"returns":  returnss,
	})
}

func GetReturnsController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	returns, err := usecase.GetReturns(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get returns",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"returnss":  returns,
	})
}

// create new returns
func CreateReturnsController(c echo.Context) error {
	returns := model.Returns{}
	c.Bind(&returns)

	if err := usecase.CreateReturns(&returns); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create returns",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new returns",
		"returns":    returns,
	})
}

// delete returns by id
func DeleteReturnsController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteReturns(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete returns",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf pengembalian tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete returns",
	})
}

// update returns by id
func UpdateReturnsController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	returns := model.Returns{}
	c.Bind(&returns)
	returns.ID = uint(id)
	if err := usecase.UpdateReturns(&returns); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update returns",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf pengembalian tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update returns",
	})
}
