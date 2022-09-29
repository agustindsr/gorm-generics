package gorm

func (r repository[T]) Create(entity *T) error {
	if result := r.GormDB.Create(entity); result.Error != nil {
		return result.Error
	}

	return nil
}
