package model

type User struct {
	ID          uint   `gorm:"primaryKey" json: "id"`
	Name        string `json:"nama"`
	Address     string `json:"alamat"`
	PhoneNumber string `json:"telp"`
	Gender      string `json:"gender"`
	Password    string `json:"password"`
}
