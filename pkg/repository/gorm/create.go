package gorm

func (r repository[T]) Create(entity T) error {
	if result := r.GormDB.Model(entity).Create(&entity); result.Error != nil {
		return result.Error
	}

	return nil
}
