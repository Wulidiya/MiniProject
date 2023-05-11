package model

import (
	"time"
)

type Transaksi struct {
	ID            uint      `gorm:"primaryKey" json: "id"`
	User_id       uint      `json:"id_anggota"`
	Tool_id       uint      `json:"id_alat"`
	BorrowingDate time.Time `json:"tanggal_peminjaman"`
	ReturnDate    time.Time `json:"tanggal_pengembalian"`
	RemainderStok string    `json:"status"`
}
