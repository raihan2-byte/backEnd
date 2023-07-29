package artikel

import "time"

type CreateArtikelFormatter struct {
	ID             int `json:"id"`
	ArtikelMessage string    `json:"message"`
	FileName       string    `json:"link_file"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FormatterArtikel (artikel Artikel) CreateArtikelFormatter {
	formatter := CreateArtikelFormatter{
		ID:        artikel.ID,
		ArtikelMessage:   artikel.ArtikelMessage,
		FileName:    artikel.FileName,
		CreatedAt: artikel.CreatedAt,
		UpdatedAt: artikel.UpdatedAt,
	}
	return formatter
}

func FormatterGetArtikel(artikel []Artikel) []CreateArtikelFormatter {
	artikelGetFormatter := []CreateArtikelFormatter{}

	for _, artikels := range artikel {
		artikelsFormatter := FormatterArtikel(artikels)
		artikelGetFormatter = append(artikelGetFormatter, artikelsFormatter)
	}

	return artikelGetFormatter
}