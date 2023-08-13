package barang

import "time"

type Barang struct {
	ID        int
	NamaPaket *string
	Nama      string
	HargaAwal *int
	Harga     int
	FileName  string
	CategoryID int
	CreatedAt time.Time
	UpdatedAt time.Time
	CategoryData Category `gorm:"foreignKey:CategoryID"`
}

type Category struct {
	ID int
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}