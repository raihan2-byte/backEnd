package artikel

type CreateArtikel struct {
	ArtikelMessage string `form:"message" binding:"required"`
}

type GetArtikel struct {
	ID int `uri:"id" binding:"required"`
}
