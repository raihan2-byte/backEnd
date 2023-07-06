package karyakmpf

type InputKMPF struct {
	Judul     string `form:"judul" binding:"required"`
	Deskripsi string `form:"deskripsi" binding:"required"`
}

type GetKMPFID struct {
	ID int `uri:"id" binding:"required"`
}