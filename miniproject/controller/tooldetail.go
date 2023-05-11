package controller

import (
	"miniproject/model"
	"miniproject/repository/database"
	"miniproject/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TooldetailController interface {
	GetTooldetailcontroller(c echo.Context) error
	GetTooldetailController(c echo.Context) error
	CreateTooldetailController(c echo.Context) error
	DeleteTooldetailController(c echo.Context) error
	UpdateTooldetailController(c echo.Context) error
}

type tooldetailController struct {
	tooldetailUsecase    usecase.TooldetailUsecase
	tooldetailRepository database.TooldetailRepository
}

func NewTooldetailController(
	tooldetailUsecase usecase.TooldetailUsecase,
	tooldetailRepository database.TooldetailRepository,
) *tooldetailController {
	return &tooldetailController{
		tooldetailUsecase,
		tooldetailRepository,
	}
}

func (b *tooldetailController) GetTooldetailcontroller(c echo.Context) error {
	tools, e := b.tooldetailRepository.GetToolsdetail()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"toolsdetail": toolsdetail,
	})
}

func (b *tooldetailController) GetTooldetailController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	tooldetail, err := b.tooldetailRepository.GetTooldetail(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get tooldetail",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"toolsdetail": tooldetail,
	})
}

// create new tooldetail
func (b *tooldetailController) CreateTooldetailController(c echo.Context) error {
	tool := model.Tooldetail{}
	c.Bind(&tooldetail)

	if err := b.tooldetailUsecase.CreateTooldetail(&tooldetail); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create tooldetail",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success create new tooldetail",
		"tooldetail": tooldetail,
	})
}

// delete tooldetail by id
func (b *tooldetailController) DeleteTooldetailController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := b.tooldetailUsecase.DeleteTooldetail(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete tooldetail",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf alat tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete tooldetail",
	})
}

// update tooldetail by id
func (b *tooldetailController) UpdateTooldetailController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	tooldetail := model.Tooldetail{}
	c.Bind(&tooldetail)
	tooldetail.ID = uint(id)
	if err := b.tooldetailUsecase.UpdateTooldetail(&tooldetail); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update tooldetail",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf alat tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update tooldetail",
	})
}
