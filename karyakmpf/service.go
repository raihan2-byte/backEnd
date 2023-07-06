package karyakmpf

type Service interface {
	CreateKMPF(input InputKMPF, FileName string) (KMPF, error)
	GetAllKMPF(input int) ([]KMPF, error)
	DeleteKMPF(ID int) (KMPF, error)
	GetOneKMPF(ID int) (KMPF, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateKMPF(input InputKMPF, FileName string) (KMPF, error) {
	kmpf := KMPF{}

	kmpf.Judul = input.Judul
	kmpf.Deskripsi = input.Deskripsi
	kmpf.FileName = FileName

	newkmpf, err := s.repository.Save(kmpf)
	if err != nil {
		return newkmpf, err
	}
	return newkmpf, nil
}

func (s *service) GetAllKMPF(input int) ([]KMPF, error) {
	kmpf, err := s.repository.FindAll()
	if err != nil {
		return kmpf, err
	}
	return kmpf, nil
}

func (s *service) DeleteKMPF(ID int) (KMPF, error) {
	kmpf, err := s.repository.FindById(ID)
	if err != nil {
		return kmpf, err
	}
	newkmpf, err := s.repository.Delete(kmpf)
	if err != nil {
		return newkmpf, err
	}
	return newkmpf, nil
}

func (s *service) GetOneKMPF(ID int) (KMPF, error) {
	kmpf, err := s.repository.FindById(ID)
	if err != nil {
		return kmpf, err
	}
	return kmpf, err
}
