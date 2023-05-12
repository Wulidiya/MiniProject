package usecase

import (
	"errors"
	"fmt"
	"MiniProject/model"
	"MiniProject/repository/database"
)

func CreateReturns(returns *model.Returns) error {

	// check transaction_id cannot be empty
	if returns.Transaction_id == "" {
		return errors.New("returns transaction_id cannot be empty")
	}

	// check return_date
	if returns.Returns_date == "" {
		return errors.New("returns returns_date cannot be empty")
	}

	err := database.CreateReturns(returns)
	if err != nil {
		return err
	}

	return nil
}

func GetReturns(id uint) (returns model.Returns, err error) {
	returns, err = database.GetReturns(id)
	if err != nil {
		fmt.Println("GetReturns: Error getting returns from database")
		return
	}
	return
}

func GetListReturnss() (returnss []model.Returns, err error) {
	returnss, err = database.GetReturnss()
	if err != nil {
		fmt.Println("GetListReturnss: Error getting returnss from database")
		return
	}
	return
}

func UpdateReturns(returns *model.Returns) (err error) {
	err = database.UpdateReturns(returns)
	if err != nil {
		fmt.Println("UpdateReturns : Error updating returns, err: ", err)
		return
	}

	return
}

func DeleteReturns(id uint) (err error) {
	returns := model.Returns{}
	returns.ID = id
	err = database.DeleteReturns(&returns)
	if err != nil {
		fmt.Println("DeleteReturns : error deleting returns, err: ", err)
		return
	}

	return
}
