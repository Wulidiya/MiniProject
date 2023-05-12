package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID   uint    `json:"user_id" form:"user_id" gorm:"unique"`
	Tool_id       string `json:"Tool_id" form:"Tool_id"`
	BorrowingDate string `json:"BorrowingDate" form:"BorrowingDate"`
	ReturnDate    string `json:"ReturnDate" form:"ReturnDte"`
	Status        string `json:"Status" form:"Status"`
}
