package controller

import (
	"miniproject/model"
	"miniproject/repository/database"
	"miniproject/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReturnController interface {
	GetReturncontroller(c echo.Context) error
	GetReturnController(c echo.Context) error
	CreateReturnController(c echo.Context) error
	DeleteReturnController(c echo.Context) error
	UpdateReturnController(c echo.Context) error
}

type ReturnController struct {
	ReturnUsecase    usecase.ReturnUsecase
	ReturnRepository database.ReturnRepository
}

func NewReturnController(
	ReturnUsecase usecase.ReturnUsecase,
	ReturnRepository database.ReturnRepository,
) *ReturnController {
	return &ReturnController{
		ReturnUsecase,
		ReturnRepository,
	}
}

func (b *ReturnController) GetReturncontroller(c echo.Context) error {
	Returns, e := b.ReturnRepository.GetReturns()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"Returns": Returns,
	})
}

func (b *ReturnController) GetReturnController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	Return, err := b.ReturnRepository.GetReturn(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get Return",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"Returns": Return,
	})
}

// create new Return
func (b. *ReturnController) CreateReturnController(c echo.Context) error {
	Return := model.Return{}
	c.Bind(&Return)

	if err := b.ReturnUsecase.CreateReturn(&Return); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create Return",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Return",
		"Return":  Return,
	})
}

// delete Return by id
func (b *ReturnController) DeleteReturnController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := b.ReturnUsecase.DeleteReturn(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete tool Return",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf pengembalian tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Return",
	})
}

// update Return by id
func (b *ReturnController) UpdateReturnController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	Return := model.Return{}
	c.Bind(&Return)
	Return.ID = uint(id)
	if err := b.ReturnUsecase.UpdateReturn(&Return); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update Return",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf alat tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update Return",
	})
}
