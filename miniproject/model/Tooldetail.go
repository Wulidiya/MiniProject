package model

type DetailAlat struct {
	ID            uint   `gorm:"primaryKey" json: "id"`
	ToolName      string `json:"nama_alat"`
	SizeVolume    string `json:"ukuran_volume"`
	Merk          string `json:"merk"`
	TotalStok     uint   `json:"jumlah_stok"`
	RemainderStok uint   `json:"sisa_stok"`
}
