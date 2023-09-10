package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	CheckEmail(input CheckEmail) (bool, error)
	SaveAvatar(id int, fileLocation string) (User, error)
	GetUserById(id int) (User, error)
}

type userService struct {
	repository UserRepository
}

func NewService(repository UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) CheckEmail(input CheckEmail) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *userService) SaveAvatar(id int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	userUpdate, err := s.repository.Update(user)
	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil

}

func (s *userService) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}
	return user, nil
}
