package home

import "gorm.io/gorm"

type Repository interface {
	//create User
	Save(tag TagLineHome) (TagLineHome, error)
	FindById(ID int) (TagLineHome, error)
	FindAll() ([]TagLineHome, error)
	// CreateImage(tag tag) (tag, error)
	// FindByEmail(email string) (tag, error)
	Update(tag TagLineHome) (TagLineHome, error)
	Delete(tag TagLineHome) (TagLineHome, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// CreateImage(berita Berita) (Berita, error){

// }

func (r *repository) FindAll() ([]TagLineHome, error) {
	var tag []TagLineHome

	err := r.db.Order("id DESC").Find(&tag).Error
	if err != nil {
		return tag, err
	}
	return tag, nil
}

func (r *repository) Save(tag TagLineHome) (TagLineHome, error) {
	err := r.db.Create(&tag).Error

	if err != nil {
		return tag, err
	}
	return tag, nil
}

// func (r *repository) FindByEmail(email string) (User, error) {
// 	var user User
// 	err := r.db.Where("email = ?", email).Find(&user).Error

// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

func (r *repository) FindById(ID int) (TagLineHome, error) {
	var tag TagLineHome

	err := r.db.Where("id = ?", ID).Find(&tag).Error

	if err != nil {
		return tag, err
	}
	return tag, nil
}

func (r *repository) Update(tag TagLineHome) (TagLineHome, error) {
	err := r.db.Save(&tag).Error
	if err != nil {
		return tag, err
	}

	return tag, nil

}

func (r *repository) Delete(tag TagLineHome) (TagLineHome, error) {
	err := r.db.Delete(&tag).Error
	if err != nil {
		return tag, err
	}

	return tag, nil
}
