package user

import "golang.org/x/crypto/bcrypt"

type UserService interface {
	RegisterUser(input RegisterUserInput) (User, error)
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
