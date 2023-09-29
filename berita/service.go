package berita

type Service interface {
	CreateBerita(input CreateBerita) (Berita, error)
	GetAllBerita(input int) ([]Berita, error)
	DeleteBerita(ID int) error
	GetOneBerita(ID int) (Berita, error)
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
