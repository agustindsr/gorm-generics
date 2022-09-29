package users

import (
	"gorm-with-generics/pkg/models"
	rmodels "gorm-with-generics/pkg/repository/users/models"
)

func (r repository) Create(user *models.User) error {
	dao := rmodels.ToDAO(*user)

	return r.gormRepo.Create(dao)
}
