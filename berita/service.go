package berita

type Service interface {
	CreateBerita(input CreateBerita, FileLocation string, fileLocation2 string, fileLocation3 string) (Berita, error)
	GetAllBerita(input int) ([]Berita, error)
	DeleteBerita(ID int) (Berita, error)
	GetOneBerita(ID int) (Berita, error)
	FindByTags(ID int) ([]Berita, error)
	FindByKarya() ([]Berita, error)
	UpdateBerita(GetIdBerita GetBerita, input CreateBerita, FileName string) (Berita, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) UpdateBerita(GetIdBerita GetBerita, input CreateBerita, fileLocation string) (Berita, error) {
	berita, err := s.repository.FindById(GetIdBerita.ID)
	if err != nil {
		return berita, err
	}

	berita.JudulBerita = input.JudulBerita
	berita.JudulBerita = input.JudulBerita
	berita.BeritaMessage = input.BeritaMessage
	berita.TagsID = input.TagsID
	berita.KaryaNewsID = input.KaryaNewsID
	berita.FileName = fileLocation

	newBerita, err := s.repository.Update(berita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
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

func (s *service) CreateBerita(input CreateBerita, fileLocation string, fileLocation2 string, fileLocation3 string) (Berita, error) {
	createBerita := Berita{}

	createBerita.JudulBerita = input.JudulBerita
	createBerita.BeritaMessage = input.BeritaMessage
	createBerita.TagsID = input.TagsID
	createBerita.KaryaNewsID = input.KaryaNewsID
	createBerita.FileName = fileLocation
	createBerita.FileName2 = fileLocation2
	createBerita.FileName3 = fileLocation3

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

func (s *service) FindByKarya() ([]Berita, error) {
	berita, err := s.repository.FindByKarya([]int{1, 2, 3})
	if err != nil {
		return berita, err
	}
	return berita, nil
}
