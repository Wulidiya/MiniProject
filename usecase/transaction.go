package usecase

import (
	"errors"
	"fmt"
	"MiniProject/model"
	"MiniProject/repository/database"
)

func CreateTransaction(transaction *model.Transaction) error {

	// check tool_id
	if transaction.Tool_id == "" {
		return errors.New("transaction tool_id cannot be empty")
	}

	err := database.CreateTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}

func GetTransaction(id uint) (transaction model.Transaction, err error) {
	transaction, err = database.GetTransaction(id)
	if err != nil {
		fmt.Println("GetTransaction: Error getting transaction from database")
		return
	}
	return
}

func GetListTransactions() (transactions []model.Transaction, err error) {
	transactions, err = database.GetTransactions()
	if err != nil {
		fmt.Println("GetListTransactions: Error getting transactions from database")
		return
	}
	return
}

func UpdateTransaction(transaction *model.Transaction) (err error) {
	err = database.UpdateTransaction(transaction)
	if err != nil {
		fmt.Println("UpdateTransaction : Error updating transaction, err: ", err)
		return
	}

	return
}

func DeleteTransaction(id uint) (err error) {
	transaction := model.Transaction{}
	transaction.ID = id
	err = database.DeleteTransaction(&transaction)
	if err != nil {
		fmt.Println("DeleteTransaction : error deleting transaction, err: ", err)
		return
	}

	return
}
