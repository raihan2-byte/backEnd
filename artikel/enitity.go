package artikel

import "time"

type Artikel struct {
	ID            int
	Judul string
	ArtikelMessage string
	FileName      string
	ImageBase64   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}