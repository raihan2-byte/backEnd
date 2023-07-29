package berita

type CreateBerita struct {
	BeritaMessage string `form:"message" binding:"required"`
	TagsID        int    `form:"tags_id" binding:"required"`
	KaryaNewsID   *int   `form:"karya_id"`
}

type GetBerita struct {
	ID int `uri:"id" binding:"required"`
}

type GetTags struct {
	Tags int `uri:"id" binding:"required"`
}

type GetKarya struct {
	Karya int `uri:"id" binding:"required"`
}