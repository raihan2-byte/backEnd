package berita

import "time"

type Berita struct {
	ID int
	BeritaMessage string
	FileName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
