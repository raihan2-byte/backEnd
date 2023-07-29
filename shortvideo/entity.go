package shortvideo

import "time"

type ShortVideo struct {
	ID        int
	Judul     string
	Deskripsi string
	FileName  string
	Source string
	CreatedAt time.Time
	UpdatedAt time.Time
}