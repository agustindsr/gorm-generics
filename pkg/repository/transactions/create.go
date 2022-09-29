package transactions

import (
	"gorm-with-generics/pkg/models"
	rmodels "gorm-with-generics/pkg/repository/transactions/models"
)

func (r repository) Create(user *models.TransactionLedger) error {
	dao := rmodels.ToDAO(*user)

	return r.gormRepo.Create(dao)
}
