package database

import (
	"miniproject/config"
	"miniproject/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *model.Transaction) error
	GetTransactions() (transactions []model.Transaction, err error)
	GetTransaction(id uint) (transaction model.Transaction, err error)
	UpdateTransaction(transaction *model.Transaction) error
	DeleteTransaction(transaction *model.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (b *transactionRepository) CreateTransaction(transaction *model.Transaction) error {
	if err := config.DB.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

func (b *transactionRepository) GetTransactions() (transactions []model.Transaction, err error) {
	if err = config.DB.Find(&transactions).Error; err != nil {
		return
	}
	return
}

func (b *transactionRepository) GetTransaction(id uint) (transaction model.Transaction, err error) {
	transaction.ID = id
	if err = config.DB.First(&transaction).Error; err != nil {
		return
	}
	return
}

func (b *transactionRepository) UpdateTransaction(transaction *model.Transaction) error {
	if err := config.DB.Updates(transaction).Error; err != nil {
		return err
	}
	return nil
}

func (b *transactionRepository) DeleteTransaction(transaction *model.Transaction) error {
	if err := config.DB.Delete(transaction).Error; err != nil {
		return err
	}
	return nil
}
