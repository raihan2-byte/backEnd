package barang

import "gorm.io/gorm"

type Repository interface {
	Save(barang Barang) (Barang, error)
	FindById(ID int) (Barang, error)
	FindAll() ([]Barang, error)
	Update(barang Barang) (Barang, error)
	Delete(barang Barang) (Barang, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(barang Barang) (Barang, error) {
	err := r.db.Create(&barang).Error

	if err != nil {
		return barang, err
	}
	return barang, nil
}

func (r *repository) FindAll() ([]Barang, error){
	var barang []Barang

	err := r.db.Find(&barang).Error

	if err != nil {
		return barang, err
	}
	return barang, nil
}


func (r *repository) FindById(ID int) (Barang, error) {
	var barang Barang

	err := r.db.Where("id = ?", ID).Find(&barang).Error

	if err != nil {
		return barang, err
	}
	return barang, nil
}

func (r *repository) Update(barang Barang) (Barang, error) {
	err := r.db.Save(&barang).Error
	if err != nil {
		return barang, err
	}

	return barang, nil

}

func (r *repository) Delete(barang Barang) (Barang, error) {
	err := r.db.Delete(&barang).Error
	if err != nil {
		return barang, err
	}

	return barang, nil
}
