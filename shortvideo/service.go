package shortvideo

type Service interface {
	CreateShortVideo(input InputShortVideo, FileName string) (ShortVideo, error)
	GetAllShortVideo(input int) ([]ShortVideo, error)
	DeleteShortVideo(ID int) (ShortVideo, error)
	GetOneShortVideo(ID int) (ShortVideo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateShortVideo(input InputShortVideo, FileName string) (ShortVideo, error) {
	shortVideo := ShortVideo{}

	shortVideo.Judul = input.Judul
	shortVideo.Deskripsi = input.Deskripsi
	shortVideo.Source = input.Source
	shortVideo.FileName = FileName

	newShortVideo, err := s.repository.Save(shortVideo)
	if err != nil {
		return newShortVideo, err
	}
	return newShortVideo, nil
}

func (s *service) GetAllShortVideo(input int) ([]ShortVideo, error) {
	shortVideo, err := s.repository.FindAll()
	if err != nil {
		return shortVideo, err
	}
	return shortVideo, nil
}

func (s *service) DeleteShortVideo(ID int) (ShortVideo, error) {
	shortVideo, err := s.repository.FindById(ID)
	if err != nil {
		return shortVideo, err
	}
	newShortVideo, err := s.repository.Delete(shortVideo)
	if err != nil {
		return newShortVideo, err
	}
	return newShortVideo, nil
}

func (s *service) GetOneShortVideo(ID int) (ShortVideo, error) {
	shortVideo, err := s.repository.FindById(ID)
	if err != nil {
		return shortVideo, err
	}
	return shortVideo, err
}
