package shortvideo

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]ShortVideo, error)
	FindById(ID int) (ShortVideo, error)
	Update(shortvideo ShortVideo) (ShortVideo, error)
	Save(shortvideo ShortVideo) (ShortVideo, error)
	Delete(shortvideo ShortVideo) (ShortVideo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]ShortVideo, error) {
	var shortvideo []ShortVideo

	err := r.db.Find(&shortvideo).Error
	if err != nil {
		return shortvideo, err
	}
	return shortvideo, nil
}

func (r *repository) Save(shortvideo ShortVideo) (ShortVideo, error) {
	err := r.db.Create(&shortvideo).Error

	if err != nil {
		return shortvideo, err
	}
	return shortvideo, nil
}

func (r *repository) FindById(ID int) (ShortVideo, error) {
	var shortvideo ShortVideo

	err := r.db.Where("id = ?", ID).Find(&shortvideo).Error

	if err != nil {
		return shortvideo, err
	}
	return shortvideo, nil
}

func (r *repository) Update(shortvideo ShortVideo) (ShortVideo, error) {
	err := r.db.Save(&shortvideo).Error
	if err != nil {
		return shortvideo, err
	}

	return shortvideo, nil

}

func (r *repository) Delete(shortvideo ShortVideo) (ShortVideo, error) {
	err := r.db.Delete(&shortvideo).Error
	if err != nil {
		return shortvideo, err
	}

	return shortvideo, nil
}