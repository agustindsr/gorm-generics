package users

import (
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
	rmodel "gorm-with-generics/pkg/repository/users/models"
)

func (r repository) Search(pagination *common.PaginationOptions, filters ...common.FilterOptions[rmodel.DAO]) ([]models.User, error) {
	results, err := r.gormRepo.Search(pagination, filters...)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for _, dao := range results {
		users = append(users, rmodel.ToModel(dao))
	}

	return users, nil
}

func WithFirstName(firstName string) common.FilterOptions[rmodel.DAO] {
	return func(dao *rmodel.DAO) {
		dao.FirstName = firstName
	}
}

func WithLastName(lastName string) common.FilterOptions[rmodel.DAO] {
	return func(dao *rmodel.DAO) {
		dao.LastName = lastName
	}
}
