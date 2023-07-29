package berita

type Service interface {
	CreateBerita(input CreateBerita, FileLocation string) (Berita, error)
	GetAllBerita(input int) ([]Berita, error)
	DeleteBerita(ID int) (Berita, error)
	GetOneBerita(ID int) (Berita, error)
	FindByTags(ID int) ([]Berita, error)
	FindByKarya(ID int) ([]Berita, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetOneBerita(ID int) (Berita, error) {
	berita, err := s.repository.FindById(ID)
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) DeleteBerita(ID int) (Berita, error) {
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

func (s *service) GetAllBerita(input int) ([]Berita, error) {
	berita, err := s.repository.FindAll()
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) CreateBerita(input CreateBerita, fileLocation string) (Berita, error) {
	createBerita := Berita{}

	createBerita.BeritaMessage = input.BeritaMessage
	createBerita.TagsID = input.TagsID
	createBerita.KaryaNewsID = input.KaryaNewsID

	createBerita.FileName = fileLocation

	newBerita, err := s.repository.Save(createBerita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}

func (s *service) FindByTags(ID int) ([]Berita, error) {
	berita, err := s.repository.FindByTags(ID)
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) FindByKarya(ID int) ([]Berita, error) {
	berita, err := s.repository.FindByKarya(ID)
	if err != nil {
		return berita, err
	}
	return berita, nil
}
