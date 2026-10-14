package role

import (
	"go-hris/config/database"
)

type Service struct{}

func (role Service) Find(param AllParam) ([]View, error) {
	var listRole []View
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Table("roles").Model([]View{}).Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listRole).Error
	return listRole, err
}

func (role Service) Create(newRole Create) error {
	err := database.DB.Table("roles").Model(Create{}).Create(&newRole).Error
	return err
}

func (role Service) Update(id string, newRole Update) error {
	err := database.DB.Table("roles").Model(Update{}).Where("id = ?", id).Updates(&newRole).Error
	return err
}

func (role Service) Erase(id string) error {
	err := database.DB.Table("roles").Delete(View{}, id).Error
	return err
}
