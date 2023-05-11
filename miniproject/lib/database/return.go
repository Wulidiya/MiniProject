package database

import (
	"miniproject/config"
	"miniproject/model"

	"gorm.io/gorm"
)

type ReturnRepository interface {
	CreateReturn(Return *model.Return) error
	GetReturns() (Returns []model.Return, err error)
	GetReturn(id uint) (Return model.Return, err error)
	UpdateReturn(Return *model.Return) error
	DeleteReturn(Return *model.Return) error
}

type ReturnRepository struct {
	db *gorm.DB
}

func NewReturnRepository(db *gorm.DB) *ReturnRepository {
	return &ReturnRepository{db}
}

func (b *ReturnRepository) CreateReturn(Return *model.Return) error {
	if err := config.DB.Create(Return).Error; err != nil {
		return err
	}
	return nil
}

func (b *ReturnRepository) GetReturns() (Returns []model.Return, err error) {
	if err = config.DB.Find(&Returns).Error; err != nil {
		return
	}
	return
}

func (b *ReturnRepository) GetReturn(id uint) (Return model.Return, err error) {
	Return.ID = id
	if err = config.DB.First(&Return).Error; err != nil {
		return
	}
	return
}

func (b *ReturnRepository) UpdateReturn(Return *model.Return) error {
	if err := config.DB.Updates(Return).Error; err != nil {
		return err
	}
	return nil
}

func (b *ReturnRepository) DeleteReturn(Return *model.Return) error {
	if err := config.DB.Delete(Return).Error; err != nil {
		return err
	}
	return nil
}
