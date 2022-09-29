package transaction

import (
	rtransaction "gorm-with-generics/pkg/repository/transactions"
	transactions "gorm-with-generics/pkg/service/trasactions"
)

type handler struct {
	Service transactions.Service
}

func New() handler {
	transactionRepo := rtransaction.NewRepository()

	return handler{
		Service: transactions.NewService(transactionRepo),
	}
}
