package berita

import "time"

type Berita struct {
	ID int
	JudulBerita string
	BeritaMessage string
	FileName   string
	TagsID int
	KaryaNewsID *int 
	CreatedAt  time.Time
	UpdatedAt  time.Time
	TagsData TagsBerita `gorm:"foreignKey:TagsID"`
	KaryaNewsData KaryaBerita `gorm:"foreignKey:KaryaNewsID"`
	
}

type TagsBerita struct {
	ID int
	Name string
	KaryaBeritaID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	KaryaBeritaData KaryaBerita `gorm:"foreignKey:KaryaBeritaID"`
}

type KaryaBerita struct{
	ID int
	Name string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}