package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByid(ID int) (User, error)
	DeleteUser(ID int) (User, error)
	GetAllUser() ([]User,error)
	IsEmaillAvailabilty(input CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) IsEmaillAvailabilty(input CheckEmailInput) (bool, error) {
	//karena hanya email maka di inisiasi hanya email
	email := input.Email

	//pengambilan algoritmanya repository yaitu findbyemail
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// ini nilainya true karena misal kita input email ini sama ga dengan email yang terdaftar dg id sekian
	//kalau g ada maka balikkanya 0 sehingga bisa di daftrakan atau availabilty
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Role = 2

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User not found that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) DeleteUser(userID int) (User, error) {
	user, err := s.repository.FindById(userID)
	if err != nil {
		return user, err
	}
	userDel, err := s.repository.Delete(user)

	if err != nil {
		return userDel, err
	}
	return userDel, nil
}

func (s *service) GetUserByid(ID int) (User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found With That ID")
	}

	return user, nil

}

func (s *service) GetAllUser() ([]User,error){

	newUser, err := s.repository.FindAll()
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
