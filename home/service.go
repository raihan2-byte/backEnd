package home

type Service interface {
	CreateTagHome(input CreateTagHome) (TagLineHome, error)
	GetAllTagLineHome(input int) ([]TagLineHome, error)
	DeleteTagLineHome(ID int) (TagLineHome, error)
	GetOneTagLineHome(ID int) (TagLineHome, error)
	UpdateTagHome(ID int, input CreateTagHome) (TagLineHome, error)
	// UpdatedUser(getUpdatedInput DeletedUser, inputUser UpdatedUser) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) UpdateTagHome(ID int, input CreateTagHome) (TagLineHome, error) {

	createTagLineHome := TagLineHome{}

	berita, err := s.repository.FindById(createTagLineHome.ID)
	if err != nil {
		return berita, err
	}

	createTagLineHome.Heading = input.Heading
	createTagLineHome.Text = input.Text

	newTagLineHome, err := s.repository.Update(createTagLineHome)
	if err != nil {
		return newTagLineHome, err
	}
	return newTagLineHome, nil
}

func (s *service) GetOneTagLineHome(ID int) (TagLineHome, error) {
	merch, err := s.repository.FindById(ID)
	if err != nil {
		return merch, err
	}
	return merch, nil
}

func (s *service) DeleteTagLineHome(ID int) (TagLineHome, error) {
	merch, err := s.repository.FindById(ID)
	if err != nil {
		return merch, err
	}

	newTagLineHome, err := s.repository.Delete(merch)
	if err != nil {
		return newTagLineHome, err
	}
	return newTagLineHome, nil
}

func (s *service) GetAllTagLineHome(input int) ([]TagLineHome, error) {
	merch, err := s.repository.FindAll()
	if err != nil {
		return merch, err
	}
	return merch, nil
}

func (s *service) CreateTagHome(input CreateTagHome) (TagLineHome, error) {
	createTagLineHome := TagLineHome{}

	createTagLineHome.Heading = input.Heading
	createTagLineHome.Text = input.Text

	newTagLineHome, err := s.repository.Save(createTagLineHome)
	if err != nil {
		return newTagLineHome, err
	}
	return newTagLineHome, nil
}
