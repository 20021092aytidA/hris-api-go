package userdetail

import (
	"go-hris/config/database"
)

type Service struct{}

func (userDetail Service) Find(param AllParam) ([]ViewModel, error) {
	var listUserDetail []ViewModel
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Preload("User").Preload("User.Role").Table("user_details").Model([]ViewModel{}).Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listUserDetail).Error
	return listUserDetail, err
}

func (userDetail Service) Create(newUserDetail CreateModel) error {
	err := database.DB.Table("user_details").Model(CreateModel{}).Create(&newUserDetail).Error
	return err
}

func (userDetail Service) Update(id string, newUserDetail UpdateModel) error {
	err := database.DB.Table("user_details").Model(UpdateModel{}).Where("id = ?", id).Updates(&newUserDetail).Error
	return err
}

func (userDetail Service) Erase(id string) error {
	err := database.DB.Table("user_details").Delete(ViewModel{}, id).Error
	return err
}
