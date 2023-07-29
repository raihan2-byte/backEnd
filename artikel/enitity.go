package artikel

import "time"

type Artikel struct {
	ID            int
	ArtikelMessage string
	FileName      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}