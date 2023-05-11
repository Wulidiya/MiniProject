package usecase

import (
	"errors"
	"fmt"
	"miniproject/model"
	"miniproject/database"
)

type ToolUsecase interface {
	CreateTool(Tool *model.Tool) error
	GetTool(id uint) (Tool model.Tool, err error)
	GetListTools() (Tools []model.Tool, err error)
	UpdateTool(Tool *model.Tool) (err error)
	DeleteTool(id uint) (err error)
}

type ToolUsecase struct {
	ToolRepository database.ToolRepository
}

func NewToolUsecase(ToolRepo database.ToolRepository) *ToolUsecase {
	return &ToolUsecase{ToolRepository: ToolRepo}
}

func (b *ToolUsecase) CreateTool(Tool *model.Tool) error {

	// check tool_id cannot be empty
	if Tool.Tool_id == "" {
		return errors.New("tool id cannot be empty")
	}

	// check tool_code
	if tool.tool_code == "" {
		return errors.New("tool creator cannot be empty")
	}

	err := b.toolRepository.Createtool(tool)
	if err != nil {
		return err
	}

	return nil
}

func (b *toolUsecase) Gettool(id uint) (tool model.tool, err error) {
	tool, err = b.toolRepository.Gettool(id)
	if err != nil {
		fmt.Println("Gettool: Error getting tool from database")
		return
	}
	return
}

func (b *toolUsecase) GetListtools() (tools []model.tool, err error) {
	tools, err = b.toolRepository.Gettools()
	if err != nil {
		fmt.Println("GetListtools: Error getting tools from database")
		return
	}
	return
}

func (b *toolUsecase) Updatetool(tool *model.tool) (err error) {
	err = b.toolRepository.Updatetool(tool)
	if err != nil {
		fmt.Println("Updatetool : Error updating tool, err: ", err)
		return
	}

	return
}

func (b *toolUsecase) Deletetool(id uint) (err error) {
	tool := model.tool{}
	tool.ID = id
	err = b.toolRepository.Deletetool(&tool)
	if err != nil {
		fmt.Println("Deletetool : error deleting tool, err: ", err)
		return
	}

	return
}