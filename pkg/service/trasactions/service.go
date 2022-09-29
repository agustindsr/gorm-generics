package transactions

import (
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
	"gorm-with-generics/pkg/repository/gorm"
	"gorm-with-generics/pkg/repository/transactions"
)

type Service interface {
	CreateTransaction(user *models.TransactionLedger) error
	SearchTransaction(options common.SearchOptions[models.TransactionLedger]) ([]models.TransactionLedger, error)
}

type service struct {
	transactionsRepository gorm.Repository[models.TransactionLedger]
}

func NewService(transactionsRepository transactions.Repository) Service {
	return service{transactionsRepository}
}

func (s service) CreateTransaction(user *models.TransactionLedger) error {
	return s.transactionsRepository.Create(user)
}

func (s service) SearchTransaction(options common.SearchOptions[models.TransactionLedger]) ([]models.TransactionLedger, error) {
	return s.transactionsRepository.Search(options)
}
