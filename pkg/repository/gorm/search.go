package gorm

import (
	"gorm-with-generics/pkg/common"
	"gorm.io/gorm"
)

func (r repository[T]) Search(pagination *common.PaginationOptions, filters ...common.FilterOptions[T]) ([]T, error) {
	results := new([]T)

	r.GormDB = buildFilters[T](r.GormDB, filters...)
	r.GormDB = buildPagination(r.GormDB, pagination)

	if resp := r.GormDB.Find(results); resp.Error != nil {
		return nil, resp.Error
	}

	return *results, nil
}

func buildPagination(db *gorm.DB, paginationOptions *common.PaginationOptions) *gorm.DB {
	if paginationOptions == nil {
		return db.Limit(common.DefaultLimit)
	}

	if paginationOptions.Limit != nil {
		db = db.Limit(*paginationOptions.Limit)
	}
	if paginationOptions.Offset != nil {
		db = db.Offset(*paginationOptions.Offset)
	}
	return db
}

func buildFilters[T any](db *gorm.DB, filterOptions ...common.FilterOptions[T]) *gorm.DB {
	var filter T
	for _, opt := range filterOptions {
		opt(&filter)
	}
	return db.Where(filter)
}
