package transactions

import (
	"fmt"
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
	rgorm "gorm-with-generics/pkg/repository/gorm"
	rmodels "gorm-with-generics/pkg/repository/transactions/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Repository interface {
	Create(user *models.TransactionLedger) error
	Search(pagination *common.PaginationOptions, filters ...common.FilterOptions[rmodels.DAO]) ([]models.TransactionLedger, error)
}

type repository struct {
	gormRepo rgorm.Repository[rmodels.DAO]
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

	db.AutoMigrate(&rmodels.DAO{})
	db.AutoMigrate(&rmodels.PackageApplicationDAO{})
	db.AutoMigrate(&rmodels.CouponApplicationDAO{})
	db.AutoMigrate(&rmodels.BalanceApplicationDAO{})

	return repository{
		gormRepo: rgorm.NewRepository[rmodels.DAO](db),
	}
}
