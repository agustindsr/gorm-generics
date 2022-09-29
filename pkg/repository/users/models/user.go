package models

import "gorm-with-generics/pkg/models"

type DAO struct {
	Id        int    `gorm:"primaryKey"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
}

func ToModel(dao DAO) models.User {
	return models.User{
		Id:        dao.Id,
		FirstName: dao.FirstName,
		LastName:  dao.LastName,
	}
}

func ToDAO(model models.User) DAO {
	return DAO{
		Id:        model.Id,
		FirstName: model.FirstName,
		LastName:  model.LastName,
	}
}

// TableName overrides the table name used by User to `profiles`
func (DAO) TableName() string {
	return "users"
}
