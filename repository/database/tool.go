package database

import (
	"MiniProject/config"
	"MiniProject/model"
)

func CreateTool(tool *model.Tool) error {
	if err := config.DB.Create(tool).Error; err != nil {
		return err
	}
	return nil
}

func GetTools() (tools []model.Tool, err error) {
	if err = config.DB.Find(&tools).Error; err != nil {
		return
	}
	return
}

func GetTool(id uint) (tool model.Tool, err error) {
	tool.ID = id
	if err = config.DB.First(&tool).Error; err != nil {
		return
	}
	return
}

func UpdateTool(tool *model.Tool) error {
	if err := config.DB.Updates(tool).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTool(tool *model.Tool) error {
	if err := config.DB.Delete(tool).Error; err != nil {
		return err
	}
	return nil
}
