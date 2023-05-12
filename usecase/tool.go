package usecase

import (
	"errors"
	"fmt"
	"MiniProject/model"
	"MiniProject/repository/database"
)

func CreateTool(tool *model.Tool) error {

	// check name cannot be empty
	if tool.ToolName == "" {
		return errors.New("tool name cannot be empty")
	}

	err := database.CreateTool(tool)
	if err != nil {
		return err
	}

	return nil
}

func GetTool(id uint) (tool model.Tool, err error) {
	tool, err = database.GetTool(id)
	if err != nil {
		fmt.Println("GetTool: Error getting tool from database")
		return
	}
	return
}

func GetListTools() (tools []model.Tool, err error) {
	tools, err = database.GetTools()
	if err != nil {
		fmt.Println("GetListTools: Error getting tools from database")
		return
	}
	return
}

func UpdateTool(tool *model.Tool) (err error) {
	err = database.UpdateTool(tool)
	if err != nil {
		fmt.Println("UpdateTool : Error updating tool, err: ", err)
		return
	}

	return
}

func DeleteTool(id uint) (err error) {
	tool := model.Tool{}
	tool.ID = id
	err = database.DeleteTool(&tool)
	if err != nil {
		fmt.Println("DeleteTool : error deleting tool, err: ", err)
		return
	}

	return
}
