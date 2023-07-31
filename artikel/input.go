package artikel

type CreateArtikel struct {
	ArtikelMessage string `form:"message" binding:"required"`
	Judul          string `form:"judul" binding:"required"`
}

type GetArtikel struct {
	ID int `uri:"id" binding:"required"`
}
