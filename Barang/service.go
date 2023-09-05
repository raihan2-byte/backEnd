package barang

type Service interface {
	CreateBarang(input InputBarang, FileLocation string) (Barang, error)
	GetAllBarang(input int) ([]Barang, error)
	DeleteBarang(ID int) (Barang, error)
	GetOneBarang(ID int) (Barang, error)
	GetBarangByCategory(category int) ([]Barang, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateBarang(input InputBarang, FileLocation string) (Barang, error) {
	createBarang := Barang{}

	createBarang.NamaPaket = input.NamaPaket
	createBarang.Nama = input.Nama
	createBarang.HargaAwal = input.HargaAwal
	createBarang.Harga = input.Harga
	createBarang.Diskon = input.Diskon
	createBarang.FileName = FileLocation
	createBarang.CategoryID = input.CategoryID

	newBarang, err := s.repository.Save(createBarang)
	if err != nil {
		return newBarang, err
	}
	return newBarang, nil
}

func (s *service) GetAllBarang(input int) ([]Barang, error) {
	barang, err := s.repository.FindAll()
	if err != nil {
		return barang, err
	}
	return barang, nil
}

func (s *service) DeleteBarang(ID int) (Barang, error) {
	barang, err := s.repository.FindById(ID)
	if err != nil {
		return barang, err
	}
	newBarang, err := s.repository.Delete(barang)
	if err != nil {
		return newBarang, err
	}
	return newBarang, nil
}

func (s *service) GetOneBarang(ID int) (Barang, error) {
	barang, err := s.repository.FindById(ID)
	if err != nil {
		return barang, err
	}
	return barang, nil
}

func (s *service) GetBarangByCategory(category int) ([]Barang, error) {
	// createBarang := Barang{}

	// createBarang.CategoryID = category.Category

	barang, err := s.repository.FindByCategory(category)
	if err != nil {
		return barang, err
	}
	return barang, nil
}
