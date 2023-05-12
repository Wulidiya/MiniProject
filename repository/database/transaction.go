package database

import (
	"MiniProject/config"
	"MiniProject/model"
)

func CreateTransaction(transaction *model.Transaction) error {
	if err := config.DB.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

func GetTransactions() (transactions []model.Transaction, err error) {
	if err = config.DB.Find(&transactions).Error; err != nil {
		return
	}
	return
}

func GetTransaction(id uint) (transaction model.Transaction, err error) {
	transaction.ID = id
	if err = config.DB.First(&transaction).Error; err != nil {
		return
	}
	return
}

func GetTransactionsByUserId(userID uint) (transaction model.Transaction, err error) {
	transaction.UserID = userID
	if err = config.DB.Find(&transaction).Error; err != nil {
		return
	}
	return
}

func UpdateTransaction(transaction *model.Transaction) error {
	if err := config.DB.Updates(transaction).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTransaction(transaction *model.Transaction) error {
	if err := config.DB.Delete(transaction).Error; err != nil {
		return err
	}
	return nil
}
