package phototalk

type InputPhoto struct {
	Judul     string `form:"judul" binding:"required"`
	Link      string `form:"link" binding:"required"`
	Deskripsi string `form:"deskripsi" binding:"required"`
}

type GetPhotoTalkID struct {
	ID int `uri:"id" binding:"required"`
}