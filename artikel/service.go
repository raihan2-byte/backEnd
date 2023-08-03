package artikel

type Service interface {
	CreateArtikel(input CreateArtikel, fileLocation string) (Artikel, error)
	GetAllArtikel(input int) ([]Artikel, error)
	DeleteArtikel(ID int) (Artikel, error)
	GetOneArtikel(ID int) (Artikel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetOneArtikel(ID int) (Artikel, error) {
	berita, err := s.repository.FindById(ID)
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) DeleteArtikel(ID int) (Artikel, error) {
	berita, err := s.repository.FindById(ID)
	if err != nil {
		return berita, err
	}

	newBerita, err := s.repository.Delete(berita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}

func (s *service) GetAllArtikel(input int) ([]Artikel, error) {
	berita, err := s.repository.FindAll()
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) CreateArtikel(input CreateArtikel, fileLocation string) (Artikel, error) {
	createBerita := Artikel{}

	createBerita.Judul = input.Judul
	createBerita.ArtikelMessage = input.ArtikelMessage
	createBerita.FileName = fileLocation

	newBerita, err := s.repository.Save(createBerita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}