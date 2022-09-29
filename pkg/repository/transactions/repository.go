package transactions

import (
	"fmt"
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
	rgorm "gorm-with-generics/pkg/repository/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Repository interface {
	Create(user *models.TransactionLedger) error
	Search(options common.SearchOptions[models.TransactionLedger]) ([]models.TransactionLedger, error)
}

type repository struct {
	gormRepo rgorm.Repository[transactionDAO]
}

func NewRepository() Repository {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&transactionDAO{})
	db.AutoMigrate(&PaymentPackageDAO{})

	return repository{
		gormRepo: rgorm.NewRepository[transactionDAO](db),
	}
}
