package shortvideo

type InputShortVideo struct {
	Judul     string `form:"judul" binding:"required"`
	Deskripsi string `form:"deskripsi" binding:"required"`
	Source    string `form:"source_link" binding:"required"`
}

type GetShortVideoID struct {
	ID int `uri:"id" binding:"required"`
}