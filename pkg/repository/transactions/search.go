package transactions

import (
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
)

func (r repository) Search(options common.SearchOptions[models.TransactionLedger]) ([]models.TransactionLedger, error) {
	searchOptionsDAO := common.SearchOptions[transactionDAO]{
		Pagination: options.Pagination,
	}

	if options.Filters != nil {
		searchOptionsDAO.Filters = toDAO(*options.Filters)
	}

	results, err := r.gormRepo.Search(searchOptionsDAO)
	if err != nil {
		return nil, err
	}

	var trasactions []models.TransactionLedger
	for _, dao := range results {
		trasactions = append(trasactions, toModel(dao))
	}

	return trasactions, nil
}
