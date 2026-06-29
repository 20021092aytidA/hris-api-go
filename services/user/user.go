package user

import (
	"go-hris/config/database"
	"go-hris/models/user"
)

func FindForPassword(param user.AllParam) (user.ViewWithPass, error) {
	var uSer user.ViewWithPass
	err := database.DB.Table("users").Model(user.ViewWithPass{}).Where(param.AllowedParam).First(&uSer).Error
	return uSer, err
}

func Find(param user.AllParam) ([]user.View, error) {
	var listUsers []user.View
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Table("users").Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listUsers).Error
	return listUsers, err
}

func Create(newUser *user.Create) (*user.Create, error) {
	var err error

	err = database.DB.Table("users").Model(user.Create{}).Create(&newUser).Error
	return newUser, err
}

func Update(id string, newUser user.Update) error {
	var err error

	err = database.DB.Table("users").Where("id = ?", id).Updates(&newUser).Error
	return err
}

func Erase(id string) error {
	err := database.DB.Table("users").Delete(user.View{}, id).Error
	return err
}
