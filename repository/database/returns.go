package database

import (
	"MiniProject/config"
	"MiniProject/model"
)

func CreateReturns(returns *model.Returns) error {
	if err := config.DB.Create(returns).Error; err != nil {
		return err
	}
	return nil
}

func GetReturnss() (returnss []model.Returns, err error) {
	if err = config.DB.Find(&returnss).Error; err != nil {
		return
	}
	return
}

func GetReturns(id uint) (returns model.Returns, err error) {
	returns.ID = id
	if err = config.DB.First(&returns).Error; err != nil {
		return
	}
	return
}

func GetReturnssByUserId(userID uint) (returns model.Returns, err error) {
	returns.UserID = userID
	if err = config.DB.Find(&returns).Error; err != nil {
		return
	}
	return
}

func UpdateReturns(returns *model.Returns) error {
	if err := config.DB.Updates(returns).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReturns(returns *model.Returns) error {
	if err := config.DB.Delete(returns).Error; err != nil {
		return err
	}
	return nil
}
