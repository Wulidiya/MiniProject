package model

import (
	"gorm.io/gorm"
)

type Tool struct {
	gorm.Model
	ToolName  string `json:"ToolName" form:"ToolName"`
	TotalStok string `json:"TotalStok" form:"TotalStok"`
	ToolID    string `json:"ToolID" form:"ToolID"`
	Merk      string `json:"Merk" form:"Merk"`
	Status    string `json:"Status" form:"Status"`
}
