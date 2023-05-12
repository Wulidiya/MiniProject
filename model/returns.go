package model

import (
	"gorm.io/gorm"
)

type Returns struct {
	gorm.Model
	Transaction_id string `json:"TransactionID" form:"TransactionID"`
	Returns_date    string `json:"ReturnDate" form:"ReturnDate"`
	Status        string `json:"Status" form:"Status"`
	UserID   uint    `json:"user_id" form:"user_id" gorm:"unique"`}
