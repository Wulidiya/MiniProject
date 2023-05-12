package controller

import (
	"net/http"
	"MiniProject/model"
	"MiniProject/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetToolsController(c echo.Context) error {
	tools, e := usecase.GetListTools()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"tools":  tools,
	})
}

func GetToolController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	tool, err := usecase.GetTool(uint(id))

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
func CreateToolController(c echo.Context) error {
	tool := model.Tool{}
	c.Bind(&tool)

	if err := usecase.CreateTool(&tool); err != nil {
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
func DeleteToolController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteTool(uint(id)); err != nil {
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
func UpdateToolController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	tool := model.Tool{}
	c.Bind(&tool)
	tool.ID = uint(id)
	if err := usecase.UpdateTool(&tool); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update tool",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf tool tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update tool",
	})
}
