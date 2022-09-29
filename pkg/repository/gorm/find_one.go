package gorm

import "gorm-with-generics/pkg/common"

func (r repository[T]) FindOne(filters ...common.FilterOptions[T]) (*T, error) {
	result := new(T)

	r.GormDB = buildFilters[T](r.GormDB, filters...)

	if resp := r.GormDB.First(result); resp.Error != nil {
		return nil, resp.Error
	}

	return result, nil
}
