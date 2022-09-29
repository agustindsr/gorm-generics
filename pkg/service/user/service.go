package user

import (
	"gorm-with-generics/pkg/models"
	ruser "gorm-with-generics/pkg/repository/users"
)

type Service interface {
	CreateUser(user *models.User) error
	SearchUsersByFirstName(firstName string) ([]models.User, error)
}

type service struct {
	userRepository ruser.Repository
}

func NewService(userRepository ruser.Repository) Service {
	return service{userRepository}
}

func (s service) CreateUser(user *models.User) error {
	return s.userRepository.Create(user)
}

func (s service) SearchUsersByFirstName(firstName string) ([]models.User, error) {
	return s.userRepository.Search(nil, ruser.WithFirstName(firstName))
}
