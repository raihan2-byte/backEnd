package karyakmpf

import "time"

type KMPF struct {
	ID        int
	Judul     string
	Deskripsi string
	FileName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}