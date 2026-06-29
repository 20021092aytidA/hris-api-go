package userdetail

import (
	"go-hris/config/database"
	"go-hris/models/userdetail"
)

func Find(param userdetail.AllParam) ([]userdetail.View, error) {
	var listUserDetail []userdetail.View
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Preload("User").Preload("User.Role").Table("user_details").Model([]userdetail.View{}).Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listUserDetail).Error
	return listUserDetail, err
}

func Create(newUserDetail userdetail.Create) error {
	err := database.DB.Table("user_details").Model(userdetail.Create{}).Create(&newUserDetail).Error
	return err
}

func Update(id string, newUserDetail userdetail.Update) error {
	err := database.DB.Table("user_details").Model(userdetail.Update{}).Where("id = ?", id).Updates(&newUserDetail).Error
	return err
}

func Erase(id string) error {
	err := database.DB.Table("user_details").Delete(userdetail.View{}, id).Error
	return err
}
