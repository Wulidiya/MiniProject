package database

import (
	"miniproject/config"
	"miniproject/model"

	"gorm.io/gorm"
)

type TooldetailRepository interface {
	CreateTooldetail(tooldetail *model.Tooldetail) error
	GetToolsdetail() (toolsdetail []model.Tooldetail, err error)
	GetTooldetail(id uint) (tooldetail model.Tooldetail, err error)
	UpdateTooldetail(tooldetail *model.Tooldetail) error
	DeleteTooldetail(tooldetail *model.Tooldetail) error
}

type tooldetailRepository struct {
	db *gorm.DB
}

func NewTooldetailRepository(db *gorm.DB) *tooldetailRepository {
	return &tooldetailRepository{db}
}

func (b *tooldetailRepository) CreateTooldetail(tooldetail *model.Tooldetail) error {
	if err := config.DB.Create(tooldetail).Error; err != nil {
		return err
	}
	return nil
}

func (b *tooldetailRepository) GetToolsdetail() (toolsdetail []model.Tooldetail, err error) {
func (b *tooldetailRepository) GetToolsdetail() (toolsdetail []model.Tooldetail, err error) {
	if err = config.DB.Find(&toolsdetail).Error; err != nil {
		return
	}
	return
}

func (b *tooldetailRepository) GetTooldetail(id uint) (tooldetail model.Tooldetail, err error) {
	tooldetail.ID = id
	if err = config.DB.First(&tooldetail).Error; err != nil {
		return
	}
	return
}

func (b *tooldetailRepository) UpdateTooldetail(tooldetail *model.Tooldetail) error {
	if err := config.DB.Updates(tooldetail).Error; err != nil {
		return err
	}
	return nil
}

func (b *tooldetailRepository) DeleteTooldetail(tooldetail *model.Tooldetail) error {
	if err := config.DB.Delete(tooldetail).Error; err != nil {
		return err
	}
	return nil
}