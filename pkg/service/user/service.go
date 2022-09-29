package user

import (
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
	"gorm-with-generics/pkg/repository/gorm"
	"gorm-with-generics/pkg/repository/users"
)

type Service interface {
	CreateUser(user *models.User) error
	SearchUsers(options common.SearchOptions[models.User]) ([]models.User, error)
}

type service struct {
	userRepository gorm.Repository[models.User]
}

func NewService(userRepository users.Repository) Service {
	return service{userRepository}
}

func (s service) CreateUser(user *models.User) error {
	return s.userRepository.Create(user)
}

func (s service) SearchUsers(options common.SearchOptions[models.User]) ([]models.User, error) {
	return s.userRepository.Search(options)
}
