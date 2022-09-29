package transactions

import "gorm-with-generics/pkg/models"

func (r repository) Create(user *models.TransactionLedger) error {
	dao := toDAO(*user)

	return r.gormRepo.Create(dao)
}
