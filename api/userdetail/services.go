package userdetail

import (
	"go-hris/config/database"
)

type Service struct{}

func (userDetail Service) Find(param AllParam) ([]View, error) {
	var listUserDetail []View
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Preload("User").Preload("User.Role").Table("user_details").Model([]View{}).Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listUserDetail).Error
	return listUserDetail, err
}

func (userDetail Service) Create(newUserDetail Create) error {
	err := database.DB.Table("user_details").Model(Create{}).Create(&newUserDetail).Error
	return err
}

func (userDetail Service) Update(id string, newUserDetail Update) error {
	err := database.DB.Table("user_details").Model(Update{}).Where("id = ?", id).Updates(&newUserDetail).Error
	return err
}

func (userDetail Service) Erase(id string) error {
	err := database.DB.Table("user_details").Delete(View{}, id).Error
	return err
}
