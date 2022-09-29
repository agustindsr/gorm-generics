package users

import (
	"fmt"
	"gorm-with-generics/pkg/common"
	"gorm-with-generics/pkg/models"
	rgorm "gorm-with-generics/pkg/repository/gorm"
	rmodel "gorm-with-generics/pkg/repository/users/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Repository interface {
	Create(user *models.User) error
	Search(pagination *common.PaginationOptions, filters ...common.FilterOptions[rmodel.DAO]) ([]models.User, error)
}

type repository struct {
	gormRepo rgorm.Repository[rmodel.DAO]
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

	db.AutoMigrate(&rmodel.DAO{})

	return repository{
		gormRepo: rgorm.NewRepository[rmodel.DAO](db),
	}
}
