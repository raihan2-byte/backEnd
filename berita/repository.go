package berita

import "gorm.io/gorm"

type Repository interface {
	//create User
	Save(berita Berita) (Berita, error)
	FindById(ID int) (Berita, error)
	FindAll() ([]Berita, error)
	// CreateImage(berita Berita) (Berita, error)
	// FindByEmail(email string) (Berita, error)
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

	err := r.db.Find(&berita).Error
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (r *repository) Save(berita Berita) (Berita, error) {
	err := r.db.Create(&berita).Error

	if err != nil {
		return berita, err
	}
	return berita, nil
}

// func (r *repository) FindByEmail(email string) (User, error) {
// 	var user User
// 	err := r.db.Where("email = ?", email).Find(&user).Error

// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

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
