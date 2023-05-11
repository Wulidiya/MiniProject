package controller

import (
	"miniproject/model"
	"miniproject/repository/database"
	"miniproject/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ToolController interface {
	GetToolcontroller(c echo.Context) error
	GetToolController(c echo.Context) error
	CreateToolController(c echo.Context) error
	DeleteToolController(c echo.Context) error
	UpdateToolController(c echo.Context) error
}

type toolController struct {
	toolUsecase    usecase.ToolUsecase
	toolRepository database.ToolRepository
}

func NewToolController(
	toolUsecase usecase.ToolUsecase,
	toolRepository database.ToolRepository,
) *toolController {
	return &toolController{
		toolUsecase,
		toolRepository,
	}
}

func (b *toolController) GetToolcontroller(c echo.Context) error {
	tools, e := b.toolRepository.GetTools()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"tools":  tools,
	})
}

func (b *toolController) GetToolController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	tool, err := b.toolRepository.GetTool(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get tool",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"tools":  tool,
	})
}

// create new tool
func (b *toolController) CreateToolController(c echo.Context) error {
	tool := model.Tool{}
	c.Bind(&tool)

	if err := b.toolUsecase.CreateTool(&tool); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create tool",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new tool",
		"tool":    tool,
	})
}

// delete tool by id
func (b *toolController) DeleteToolController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := b.toolUsecase.DeleteTool(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete tool",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf alat tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete tool",
	})
}

// update tool by id
func (b *toolController) UpdateToolController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	tool := model.Tool{}
	c.Bind(&tool)
	tool.ID = uint(id)
	if err := b.toolUsecase.UpdateTool(&tool); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update tool",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf alat tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update tool",
	})
}
