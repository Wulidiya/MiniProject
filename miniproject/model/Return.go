package model

import (
	"time"
)

type Pengembalian struct {
	ID            uint      `gorm:"primaryKey" json: "id"`
	TransactionID uint      `json:"id_transaksi"`
	ReturnDate    time.Time `json:"tanggal_pengembalian"`
	Status        string    `json:"status"`
}
