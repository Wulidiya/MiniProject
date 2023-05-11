package model

type Alat struct {
	ID       uint   `gorm:"primaryKey" json: "id"`
	ToolID   uint   `json:"id_alat"`
	ToolCode string `json:"kode_alat"`
	Status   string `json:"status"`
}
