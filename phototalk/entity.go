package phototalk

import "time"

type PhotoTalk struct {
	ID        int
	Judul     string
	Deskripsi string
	FileName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}