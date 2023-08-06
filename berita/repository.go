package berita

import "gorm.io/gorm"

type Repository interface {
	//create User
	Save(berita Berita) (Berita, error)
	FindById(ID int) (Berita, error)
	FindAll() ([]Berita, error)
	FindByKarya(Karya int) ([]Berita, error)
	FindByTags(tags int) ([]Berita, error)
	Update(berita Berita) (Berita, error)
	Delete(berita Berita) (Berita, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// CreateImage(berita Berita) (Berita, error){

// }

func (r *repository) FindAll()([]Berita, error){
	var berita []Berita

	err := r.db.Order("id DESC").Preload("TagsData").Preload("KaryaNewsData").Find(&berita).Error
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

func (r *repository) FindByTags(tags int) ([]Berita, error) {
	var berita []Berita

	err := r.db.Preload("TagsData").Preload("KaryaNewsData").Where("tags_id = ?", tags).Find(&berita).Error

	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) FindByKarya(Karya int) ([]Berita, error) {
	var berita []Berita

	err := r.db.Preload("TagsData").Preload("KaryaNewsData").Where("karya_news_id = ?", Karya).Find(&berita).Error
	if err != nil {
		return berita, err
	}
	return berita, nil
}


func (r *repository) FindById(ID int) (Berita, error) {
	var berita Berita

	err := r.db.Where("id = ?", ID).Find(&berita).Error

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
