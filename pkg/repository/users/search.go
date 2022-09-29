package users

import (
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
)

func (r repository) Search(options common.SearchOptions[models.User]) ([]models.User, error) {
	searchOptionsDAO := common.SearchOptions[userDAO]{
		Filters: toDAO(*options.Filters),
		Pagination: options.Pagination,
	}

	results, err := r.gormRepo.Search(searchOptionsDAO)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for _, dao := range results {
		users = append(users, toModel(dao))
	}

	return users, nil
}
