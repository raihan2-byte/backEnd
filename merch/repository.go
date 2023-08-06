package merch

import "gorm.io/gorm"

type Repository interface {
	//create User
	Save(merch Merch) (Merch, error)
	FindById(ID int) (Merch, error)
	FindAll() ([]Merch, error)
	// CreateImage(merch merch) (merch, error)
	// FindByEmail(email string) (merch, error)
	Update(merch Merch) (Merch, error)
	Delete(merch Merch) (Merch, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// CreateImage(berita Berita) (Berita, error){

// }

func (r *repository) FindAll()([]Merch, error){
	var merch []Merch

	err := r.db.Order("id DESC").Find(&merch).Error
	if err != nil {
		return merch, err
	}
	return merch, nil
}

func (r *repository) Save(merch Merch) (Merch, error) {
	err := r.db.Create(&merch).Error

	if err != nil {
		return merch, err
	}
	return merch, nil
}

// func (r *repository) FindByEmail(email string) (User, error) {
// 	var user User
// 	err := r.db.Where("email = ?", email).Find(&user).Error

// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

func (r *repository) FindById(ID int) (Merch, error) {
	var merch Merch

	err := r.db.Where("id = ?", ID).Find(&merch).Error

	if err != nil {
		return merch, err
	}
	return merch, nil
}

func (r *repository) Update(merch Merch) (Merch, error) {
	err := r.db.Save(&merch).Error
	if err != nil {
		return merch, err
	}

	return merch, nil

}

func (r *repository) Delete(merch Merch) (Merch, error) {
	err := r.db.Delete(&merch).Error
	if err != nil {
		return merch, err
	}

	return merch, nil
}
