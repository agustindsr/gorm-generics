package users

import "gorm-with-generics/pkg/models"

func (r repository) Create(user *models.User) error {
	dao := toDAO(*user)

	return r.gormRepo.Create(dao)
}
