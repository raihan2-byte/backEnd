package phototalk

type Service interface {
	CreatePhotoTalk(input InputPhoto, FileName string) (PhotoTalk, error)
	GetAllPhotoTalk(input int) ([]PhotoTalk, error)
	DeletePhotoTalk(ID int) (PhotoTalk, error)
	GetOnePhotoTalk(ID int) (PhotoTalk, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePhotoTalk(input InputPhoto, FileName string) (PhotoTalk, error) {
	photoTalk := PhotoTalk{}

	photoTalk.Judul = input.Judul
	photoTalk.ShortDeskripsi = input.ShortDeskripsi
	photoTalk.Deskripsi = input.Deskripsi
	photoTalk.Link = input.Link
	photoTalk.FileName = FileName

	newPhotoTalk, err := s.repository.Save(photoTalk)
	if err != nil {
		return newPhotoTalk, err
	}
	return newPhotoTalk, nil
}

func (s *service) GetAllPhotoTalk(input int) ([]PhotoTalk, error) {
	photoTalk, err := s.repository.FindAll()
	if err != nil {
		return photoTalk, err
	}
	return photoTalk, nil
}

func (s *service) DeletePhotoTalk(ID int) (PhotoTalk, error) {
	photoTalk, err := s.repository.FindById(ID)
	if err != nil {
		return photoTalk, err
	}
	newPhotoTalk, err := s.repository.Delete(photoTalk)
	if err != nil {
		return newPhotoTalk, err
	}
	return newPhotoTalk, nil
}

func (s *service) GetOnePhotoTalk(ID int) (PhotoTalk, error) {
	photoTalk, err := s.repository.FindById(ID)
	if err != nil {
		return photoTalk, err
	}
	return photoTalk, err
}
