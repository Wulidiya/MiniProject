package database

import (
	"miniproject/config"
	"miniproject/model"

	"gorm.io/gorm"
)

type ToolRepository interface {
	CreateTool(tool *model.Tool) error
	GetTools() (tools []model.Tool, err error)
	GetTool(id uint) (tool model.Tool, err error)
	UpdateTool(tool *model.Tool) error
	DeleteTool(tool *model.Tool) error
}

type toolRepository struct {
	db *gorm.DB
}

func NewToolRepository(db *gorm.DB) *toolRepository {
	return &toolRepository{db}
}

func (b *toolRepository) CreateTool(tool *model.Tool) error {
	if err := config.DB.Create(tool).Error; err != nil {
		return err
	}
	return nil
}

func (b *toolRepository) GetTools() (tools []model.Tool, err error) {
	if err = config.DB.Find(&tools).Error; err != nil {
		return
	}
	return
}

func (b *toolRepository) GetTool(id uint) (tool model.Tool, err error) {
	tool.ID = id
	if err = config.DB.First(&tool).Error; err != nil {
		return
	}
	return
}

func (b *toolRepository) UpdateTool(tool *model.Tool) error {
	if err := config.DB.Updates(tool).Error; err != nil {
		return err
	}
	return nil
}

func (b *toolRepository) DeleteTool(tool *model.Tool) error {
	if err := config.DB.Delete(tool).Error; err != nil {
		return err
	}
	return nil
}