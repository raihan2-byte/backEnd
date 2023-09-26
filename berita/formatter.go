package berita

import "time"

type BeritaStructFormatter struct {
	ID            int         `json:"id"`
	Judul string      `json:"judul"`
	BeritaMessage string      `json:"berita_message"`
	TagsID        int         `json:"tags_id"`
	KaryaNewsID   *int        `json:"karya_news_id"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	FileName      []string      `json:"file_names"`
	TagsData      TagsDataStructFormatter  `gorm:"foreignKey:tags_id"`
	KaryaNewsData KaryaBeritaStructFormatter `gorm:"foreignKey:karya_news_id"`
}


type BeritaImagesFormatter struct {
    FileName  string `json:"berita_message"`
	BeritaID  int `json:"Ber"`
}

type TagsDataStructFormatter struct {
	ID int `json:"id"`
	Name string `json:"nama"`
	KaryaBeritaID int `json:"karya_berita_id"`
}

type KaryaBeritaStructFormatter struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func FormatterBerita (berita Berita) BeritaStructFormatter{
	formatter := BeritaStructFormatter{
		ID: berita.ID,
		Judul: berita.JudulBerita,
		BeritaMessage: berita.BeritaMessage,
		FileName: make([]string, len(berita.FileName)),
		TagsID: berita.TagsID,
		KaryaNewsID: berita.KaryaNewsID,
		CreatedAt: berita.CreatedAt,
		UpdatedAt: berita.UpdatedAt,
	}
	for i, image := range berita.FileName {
        formatter.FileName[i] = image.FileName
    }

	tagsData := berita.TagsData

	tagsDataFormatter := TagsDataStructFormatter{}
	tagsDataFormatter.ID = tagsData.ID
	tagsDataFormatter.Name = tagsData.Name
	tagsDataFormatter.KaryaBeritaID = tagsData.KaryaBeritaID

	formatter.TagsData = tagsDataFormatter

	karyaNewsData := berita.KaryaNewsData

	karya := KaryaBeritaStructFormatter{}
	karya.ID = karyaNewsData.ID
	karya.Name = karyaNewsData.Name

	formatter.KaryaNewsData = karya

	return formatter
}


func FormatterGetBerita(berita []Berita) []BeritaStructFormatter {
	beritaGetFormatter := []BeritaStructFormatter{}

	for _, beritas := range berita {
		beritaFormatter := FormatterBerita(beritas)
		beritaGetFormatter = append(beritaGetFormatter, beritaFormatter)
	}

	return beritaGetFormatter
}

type BeritaSingleFormatter struct {
    ID            int         `json:"id"`
    Judul         string      `json:"judul"`
    BeritaMessage string      `json:"berita_message"`
    TagsID        int         `json:"tags_id"`
    KaryaNewsID   *int        `json:"karya_news_id"`
    CreatedAt     time.Time   `json:"created_at"`
    UpdatedAt     time.Time   `json:"updated_at"`
    FileNames     []string    `json:"file_names"`
    TagsData      TagsDataStructFormatter  `gorm:"foreignKey:tags_id"`
    KaryaNewsData KaryaBeritaStructFormatter `gorm:"foreignKey:karya_news_id"`
}

func FormatterOneBerita(berita Berita) BeritaSingleFormatter {
    formatter := BeritaSingleFormatter{
        ID:            berita.ID,
        Judul:         berita.JudulBerita,
        BeritaMessage: berita.BeritaMessage,
        FileNames:     make([]string, len(berita.FileName)),
        TagsID:        berita.TagsID,
        KaryaNewsID:   berita.KaryaNewsID,
        CreatedAt:     berita.CreatedAt,
        UpdatedAt:     berita.UpdatedAt,
    }

    for i, image := range berita.FileName {
        formatter.FileNames[i] = image.FileName
    }

    tagsData := berita.TagsData
    tagsDataFormatter := TagsDataStructFormatter{}
    tagsDataFormatter.ID = tagsData.ID
    tagsDataFormatter.Name = tagsData.Name
    tagsDataFormatter.KaryaBeritaID = tagsData.KaryaBeritaID
    formatter.TagsData = tagsDataFormatter

    karyaNewsData := berita.KaryaNewsData
    karya := KaryaBeritaStructFormatter{}
    karya.ID = karyaNewsData.ID
    karya.Name = karyaNewsData.Name
    formatter.KaryaNewsData = karya

    return formatter
}
