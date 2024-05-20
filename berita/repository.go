package berita

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	//create User
	Save(berita Berita) (Berita, error)
	CreateImage(berita BeritaImage) error
	FindById(ID int) (Berita, error)
	FindAll() ([]Berita, error)
	FindByLastID() (Berita, error)
	FindBySlug(slug string) (Berita, error)
	FindByKarya(ID []int) ([]Berita, error)
	FindByTags(tags int) ([]Berita, error)
	Update(berita Berita) (Berita, error)
	Delete(berita Berita) (Berita, error)
	DeleteImages(beritaID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByLastID() (Berita, error) {
	var berita Berita
	err := r.db.Order("created_at desc").Preload("TagsData").Preload("KaryaNewsData").Preload("FileName").First(&berita).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return berita, errors.New("no berita found")
	}
	return berita, err
}

func (r *repository) FindBySlug(slug string) (Berita, error) {
	var ecopedia Berita

	err := r.db.Where("slug = ?", slug).Preload("FileName").Find(&ecopedia).Error

	if err != nil {
		return ecopedia, err
	}
	if ecopedia.Slug == "" {
		return ecopedia, errors.New("slug not found")
	}

	return ecopedia, nil

}

func (r *repository) FindAll() ([]Berita, error) {
	var berita []Berita

	err := r.db.Order("id DESC").Preload("TagsData").Preload("KaryaNewsData").Preload("FileName").Find(&berita).Error
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) Save(berita Berita) (Berita, error) {
	err := r.db.Preload("TagsData").Preload("KaryaNewsData").Create(&berita).Error

	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) CreateImage(berita BeritaImage) error {
	err := r.db.Create(&berita).Error
	return err

}

func (r *repository) FindByTags(tags int) ([]Berita, error) {
	var berita []Berita

	err := r.db.Preload("TagsData").Preload("KaryaNewsData").Where("tags_id = ?", tags).Find(&berita).Error

	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) FindByKarya(ID []int) ([]Berita, error) {
	var berita []Berita

	err := r.db.Preload("TagsData").Preload("KaryaNewsData").Where("karya_news_id IN ?", ID).Find(&berita).Error
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) FindById(ID int) (Berita, error) {
	var berita Berita

	err := r.db.Where("id = ?", ID).Preload("TagsData").Preload("KaryaNewsData").Preload("FileName").Find(&berita).Error

	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) Update(berita Berita) (Berita, error) {
	err := r.db.Save(&berita).Error
	if err != nil {
		return berita, err
	}

	return berita, nil

}

func (r *repository) Delete(berita Berita) (Berita, error) {
	err := r.db.Delete(&berita).Error
	if err != nil {
		return berita, err
	}

	return berita, nil
}

func (r *repository) DeleteImages(beritaID int) error {
	err := r.db.Where("berita_id = ?", beritaID).Delete(&BeritaImage{}).Error
	if err != nil {
		return err
	}

	return nil
}
