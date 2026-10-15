package role

import (
	"go-hris/config/database"
)

type Service struct{}

func (role Service) Find(param AllParam) ([]ViewModel, error) {
	var listRole []ViewModel
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Table("roles").Model([]ViewModel{}).Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listRole).Error
	return listRole, err
}

func (role Service) Create(newRole CreateModel) error {
	err := database.DB.Table("roles").Model(CreateModel{}).Create(&newRole).Error
	return err
}

func (role Service) Update(id string, newRole UpdateModel) error {
	err := database.DB.Table("roles").Model(UpdateModel{}).Where("id = ?", id).Updates(&newRole).Error
	return err
}

func (role Service) Erase(id string) error {
	err := database.DB.Table("roles").Delete(ViewModel{}, id).Error
	return err
}
