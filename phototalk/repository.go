package phototalk

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]PhotoTalk, error)
	FindById(ID int) (PhotoTalk, error)
	Update(photoTalk PhotoTalk) (PhotoTalk, error)
	Save(photoTalk PhotoTalk) (PhotoTalk, error)
	Delete(photoTalk PhotoTalk) (PhotoTalk, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]PhotoTalk, error){
	var photoTalk []PhotoTalk

	err := r.db.Order("id DESC").Find(&photoTalk).Error
	if err != nil {
		return photoTalk, err
	}
	return photoTalk, nil
}

func (r *repository) Save(photoTalk PhotoTalk) (PhotoTalk, error) {
	err := r.db.Create(&photoTalk).Error

	if err != nil {
		return photoTalk, err
	}
	return photoTalk, nil
}

func (r *repository) FindById(ID int) (PhotoTalk, error) {
	var photoTalk PhotoTalk

	err := r.db.Where("id = ?", ID).Find(&photoTalk).Error

	if err != nil {
		return photoTalk, err
	}
	return photoTalk, nil
}

func (r *repository) Update(photoTalk PhotoTalk) (PhotoTalk, error) {
	err := r.db.Save(&photoTalk).Error
	if err != nil {
		return photoTalk, err
	}

	return photoTalk, nil

}

func (r *repository) Delete(photoTalk PhotoTalk) (PhotoTalk, error) {
	err := r.db.Delete(&photoTalk).Error
	if err != nil {
		return photoTalk, err
	}

	return photoTalk, nil
}