package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailability(input CheckEmailInput) (bool, error)
	SaveAvatar(id int, filelocation string) (User, error)
}

type service struct {
	repo Repository
}

// SaveAvatar implements Service.
func (s *service) SaveAvatar(id int, filelocation string) (User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = filelocation

	updated, err := s.repo.UpdateUser(user)
	if err != nil {
		return updated, err
	}

	return updated, nil
}

// IsEmailAvailability implements Service.
func (s *service) IsEmailAvailability(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

// Login implements Service.
func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.HashPassword = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repo.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func NewService(repo Repository) Service {
	return &service{repo}
}
