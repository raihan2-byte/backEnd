package karyakmpf

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]KMPF, error)
	FindById(ID int) (KMPF, error)
	Update(kmpf KMPF) (KMPF, error)
	Save(kmpf KMPF) (KMPF, error)
	Delete(kmpf KMPF) (KMPF, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]KMPF, error){
	var kmpf []KMPF

	err := r.db.Find(&kmpf).Error
	if err != nil {
		return kmpf, err
	}
	return kmpf, nil
}

func (r *repository) Save(kmpf KMPF) (KMPF, error) {
	err := r.db.Create(&kmpf).Error

	if err != nil {
		return kmpf, err
	}
	return kmpf, nil
}

func (r *repository) FindById(ID int) (KMPF, error) {
	var kmpf KMPF

	err := r.db.Where("id = ?", ID).Find(&kmpf).Error

	if err != nil {
		return kmpf, err
	}
	return kmpf, nil
}

func (r *repository) Update(kmpf KMPF) (KMPF, error) {
	err := r.db.Save(&kmpf).Error
	if err != nil {
		return kmpf, err
	}

	return kmpf, nil

}

func (r *repository) Delete(kmpf KMPF) (KMPF, error) {
	err := r.db.Delete(&kmpf).Error
	if err != nil {
		return kmpf, err
	}

	return kmpf, nil
}