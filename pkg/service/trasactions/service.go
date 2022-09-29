package transactions

import (
	"gorm-with-generics/pkg/models"
	"gorm-with-generics/pkg/repository/transactions"
)

type Service interface {
	CreateTransaction(user *models.TransactionLedger) error
	SearchTransactionByAmount(amount int64) ([]models.TransactionLedger, error)
}

type service struct {
	transactionsRepository transactions.Repository
}

func NewService(transactionsRepository transactions.Repository) Service {
	return service{transactionsRepository}
}

func (s service) CreateTransaction(user *models.TransactionLedger) error {
	return s.transactionsRepository.Create(user)
}

func (s service) SearchTransactionByAmount(amount int64) ([]models.TransactionLedger, error) {
	return s.transactionsRepository.Search(nil, transactions.WithAmount(amount))
}
