package user

import (
	"go-hris/config/database"
)

type Service struct{}

func (user Service) FindForPassword(param AllParam) (ViewWithPass, error) {
	var uSer ViewWithPass
	err := database.DB.Table("users").Model(ViewWithPass{}).Where(param.AllowedParam).First(&uSer).Error
	return uSer, err
}

func (user Service) Find(param AllParam) ([]View, error) {
	var listUsers []View
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Table("users").Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listUsers).Error
	return listUsers, err
}

func (user Service) Create(newUser *Create) (*Create, error) {
	var err error

	err = database.DB.Table("users").Model(Create{}).Create(&newUser).Error
	return newUser, err
}

func (user Service) Update(id string, newUser Update) error {
	var err error

	err = database.DB.Table("users").Where("id = ?", id).Updates(&newUser).Error
	return err
}

func (user Service) Erase(id string) error {
	err := database.DB.Table("users").Delete(View{}, id).Error
	return err
}
