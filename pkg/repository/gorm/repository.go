package gorm

import (
	"gorm-with-generics/pkg/common"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(entity *T) error
	Search(options common.SearchOptions[T]) ([]T, error)
}

type repository[T any] struct {
	GormDB *gorm.DB
}

func NewRepository[T any](gormDB *gorm.DB) Repository[T] {
	return repository[T]{
		GormDB: gormDB,
	}
}
