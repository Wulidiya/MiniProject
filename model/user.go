package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name" form:"name" validate:"required,max=20"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Password    string `json:"password" form:"password" validate:"required,min=5"`
	Address     string `json:"Address" form:"Address"`
	PhoneNumber string `json:"Phone_Number" form:"Phone_Number"`
	Gender      string `json:"Gender" form:"Gender"`
	Token    string `gorm:"-"`
	UserID   uint    `json:"user_id" form:"user_id" gorm:"unique"`
}
