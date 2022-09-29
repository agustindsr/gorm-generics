package users

import "gorm-with-generics/pkg/models"

type userDAO struct {
	Id        int    `gorm:"primaryKey"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
}

func toModel(dao userDAO) models.User {
	return models.User{
		Id:        dao.Id,
		FirstName: dao.FirstName,
		LastName:  dao.LastName,
	}
}

func toDAO(model models.User) *userDAO {
	return &userDAO{
		Id:        model.Id,
		FirstName: model.FirstName,
		LastName:  model.LastName,
	}
}

// TableName overrides the table name used by User to `profiles`
func (userDAO) TableName() string {
	return "users"
}
