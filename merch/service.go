package merch

type Service interface {
	CreateMerch(input CreateMerch, FileLocation string) (Merch, error)
	GetAllMerch(input int) ([]Merch, error)
	DeleteMerch(ID int) (Merch, error)
	GetOneMerch(ID int) (Merch, error)
	// UpdatedUser(getUpdatedInput DeletedUser, inputUser UpdatedUser) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetOneMerch(ID int) (Merch, error) {
	merch, err := s.repository.FindById(ID)
	if err != nil {
		return merch, err
	}
	return merch, nil
}

func (s *service) DeleteMerch(ID int) (Merch, error) {
	merch, err := s.repository.FindById(ID)
	if err != nil {
		return merch, err
	}

	newMerch, err := s.repository.Delete(merch)
	if err != nil {
		return newMerch, err
	}
	return newMerch, nil
}

func (s *service) GetAllMerch(input int) ([]Merch, error) {
	merch, err := s.repository.FindAll()
	if err != nil {
		return merch, err
	}
	return merch, nil
}

func (s *service) CreateMerch(input CreateMerch, fileLocation string) (Merch, error) {
	createMerch := Merch{}

	createMerch.Name = input.Name
	createMerch.Price = input.Price
	createMerch.Link = input.Link
	createMerch.FileName = fileLocation

	newMerch, err := s.repository.Save(createMerch)
	if err != nil {
		return newMerch, err
	}
	return newMerch, nil
}
