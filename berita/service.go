package berita

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	GetLatestBerita() (Berita, error)
	CreateBerita(input CreateBerita) (Berita, error)
	GetAllBerita(input int) ([]Berita, error)
	DeleteBerita(ID int) error
	GetOneBerita(slug string) (Berita, error)
	FindByTags(ID int) ([]Berita, error)
	CreateBeritaImage(BeritaID int, FileName string) error
	FindByKarya() ([]Berita, error)
	UpdateBerita(GetIdBerita GetBerita, input CreateBerita) (Berita, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetLatestBerita() (Berita, error) {
	return s.repository.FindByLastID()
}

func (s *service) CreateBeritaImage(beritaID int, FileName string) error {
	createBerita := BeritaImage{}

	createBerita.FileName = FileName
	createBerita.BeritaID = beritaID

	err := s.repository.CreateImage(createBerita)
	if err != nil {
		return err
	}
	return nil

}

func (s *service) UpdateBerita(GetIdBerita GetBerita, input CreateBerita) (Berita, error) {
	berita, err := s.repository.FindById(GetIdBerita.ID)
	if err != nil {
		return berita, err
	}

	berita.JudulBerita = input.JudulBerita
	berita.JudulBerita = input.JudulBerita
	berita.BeritaMessage = input.BeritaMessage
	berita.TagsID = input.TagsID
	berita.KaryaNewsID = input.KaryaNewsID

	oldSlug := berita.Slug
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	slugTitle := strings.ToLower(input.JudulBerita)
	mySlug := slug.Make(slugTitle)
	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999
	berita.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

	// Ubah nilai slug kembali ke nilai slug lama untuk mencegah perubahan slug dalam database
	berita.Slug = oldSlug

	newBerita, err := s.repository.Update(berita)
	if err != nil {
		return newBerita, err
	}
	return newBerita, nil
}

func (s *service) GetOneBerita(slug string) (Berita, error) {
	berita, err := s.repository.FindBySlug(slug)
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) DeleteBerita(ID int) error {
	// Konversi beritaID menjadi tipe data yang sesuai (misalnya, int)
	// Lakukan validasi ID jika diperlukan

	// Cek apakah berita ada di basis data
	berita, err := s.repository.FindById(ID)
	if err != nil {
		return err // Berita tidak ditemukan atau terjadi kesalahan lainnya
	}

	news, err := s.repository.Delete(berita)
	if err != nil {
		return err // Tangani kesalahan jika penghapusan berita gagal
	}

	// Hapus gambar yang terkait dengan berita jika ada
	err = s.repository.DeleteImages(news.ID)
	if err != nil {
		return err // Tangani kesalahan jika penghapusan gambar gagal
	}

	// Hapus berita dari basis data

	return nil
}

func (s *service) GetAllBerita(input int) ([]Berita, error) {
	berita, err := s.repository.FindAll()
	if err != nil {
		return berita, err
	}
	return berita, nil
}

func (s *service) CreateBerita(input CreateBerita) (Berita, error) {
	createBerita := Berita{}
	createBerita.JudulBerita = input.JudulBerita
	createBerita.BeritaMessage = input.BeritaMessage
	createBerita.TagsID = input.TagsID
	createBerita.KaryaNewsID = input.KaryaNewsID

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	slugTitle := strings.ToLower(createBerita.JudulBerita)

	mySlug := slug.Make(slugTitle)

	randomNumber := seededRand.Intn(1000000) // Angka acak 0-999999

	createBerita.Slug = fmt.Sprintf("%s-%d", mySlug, randomNumber)

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
