package barang

import "time"

type CreateBarangFormatter struct {
	ID        int `json:"id"`
	NamaPaket    *string `json:"nama_paket"`
	Nama      string `json:"nama"`
	HargaAwal    *int    `json:"harga_awal"`
	Harga     int `json:"harga"`
	FileName  string `json:"file_name"`
	CategoryID int `json:"category_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Category CategoryBarang `json:"category_barang"`
}

type CategoryBarang struct {
	ID int `json:"id"`
	Name string `json:"nama"`
} 

func FormatterBarang (barang Barang) CreateBarangFormatter {
	formatter := CreateBarangFormatter{
		ID:        barang.ID,
		NamaPaket: barang.NamaPaket,
		Nama:   barang.Nama,
		HargaAwal: barang.HargaAwal,
		Harga: barang.Harga,
		FileName:    barang.FileName,
		CategoryID : barang.CategoryID,
		CreatedAt: barang.CreatedAt,
		UpdatedAt: barang.UpdatedAt,
	}

	category := barang.CategoryData

	categoryFormatter := CategoryBarang{}
	categoryFormatter.ID = category.ID
	categoryFormatter.Name = category.Name

	formatter.Category = categoryFormatter
	
	return formatter
}

func FormatterGetCategory(category []Barang) []CreateBarangFormatter {
	barangGetFormatter := []CreateBarangFormatter{}

	for _, barang := range category {
		barangFormatter := FormatterBarang(barang)
		barangGetFormatter = append(barangGetFormatter, barangFormatter)
	}

	return barangGetFormatter
}