package berita

type CreateBerita struct {
	BeritaMessage string `form:"message" binding:"required"`
}

type GetBerita struct {
	ID int `uri:"id" binding:"required"`
}
