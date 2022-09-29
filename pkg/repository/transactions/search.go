package transactions

import (
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
	rmodels "gorm-with-generics/pkg/repository/transactions/models"
)

func (r repository) Search(pagination *common.PaginationOptions, filters ...common.FilterOptions[rmodels.DAO]) ([]models.TransactionLedger, error) {
	results, err := r.gormRepo.Search(pagination, filters...)
	if err != nil {
		return nil, err
	}

	var trasactions []models.TransactionLedger
	for _, dao := range results {
		trasactions = append(trasactions, rmodels.ToModel(dao))
	}

	return trasactions, nil
}

func WithAmount(amount int64) common.FilterOptions[rmodels.DAO] {
	return func(dao *rmodels.DAO) {
		dao.Amount = amount
	}
}
