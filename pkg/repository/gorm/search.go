package gorm

import (
	"gorm-with-generics/pkg/common"
	"gorm.io/gorm"
)

func (r repository[T]) Search(options common.SearchOptions[T]) ([]T, error) {
	results := new([]T)

	r.GormDB = buildSearchOptions[T](r.GormDB, options)

	if resp := r.GormDB.Find(results); resp.Error != nil {
		return nil, resp.Error
	}

	return *results, nil
}

func buildSearchOptions[T any](db *gorm.DB, options common.SearchOptions[T]) *gorm.DB {
	if options.Filters != nil {
		db = buildFilter(db, options.Filters)
	}

	if options.Pagination != nil {
		db = buildPagination(db, options.Pagination)
	}

	return db
}

func buildPagination(db *gorm.DB, paginationOptions *common.PaginationOptions) *gorm.DB {
	return db.Offset(paginationOptions.Offset).Limit(paginationOptions.Limit)
}

func buildFilter[T any](db *gorm.DB, filterObject T) *gorm.DB {
	return db.Where(filterObject)
}
