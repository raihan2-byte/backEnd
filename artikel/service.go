package artikel

import (
	"encoding/base64"
	"os"
	"path/filepath"
)

type Service interface {
	CreateArtikel(input CreateArtikel) (Artikel, error)
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

func (s *service) CreateArtikel(input CreateArtikel) (Artikel, error) {
	createBerita := Artikel{}

	imageData, err := base64.StdEncoding.DecodeString(input.ImageBase64)
	if err != nil {
		return Artikel{}, err
	}

	fileLocation := "images" 
	err = os.MkdirAll(fileLocation, 0755)
    if err != nil {
        return Artikel{}, err
    }

	fileLocation = filepath.Join(fileLocation, input.Judul+".png")
	err = os.WriteFile(fileLocation, imageData, 0644)
	if err != nil {
		return Artikel{}, err
	}

	createBerita.Judul = input.Judul
	createBerita.ArtikelMessage = input.ArtikelMessage
	createBerita.FileName = fileLocation

	newBerita, err := s.repository.Save(createBerita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}