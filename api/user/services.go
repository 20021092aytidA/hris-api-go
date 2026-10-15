package user

import (
	"go-hris/config/database"
)

type Service struct{}

func (user Service) FindForPassword(param AllParam) (ViewWithPassModel, error) {
	var uSer ViewWithPassModel
	err := database.DB.Table("users").Model(ViewWithPassModel{}).Where(param.AllowedParam).First(&uSer).Error
	return uSer, err
}

func (user Service) Find(param AllParam) ([]ViewModel, error) {
	var listUsers []ViewModel
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Table("users").Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listUsers).Error
	return listUsers, err
}

func (user Service) Create(newUser *CreateModel) (*CreateModel, error) {
	var err error

	err = database.DB.Table("users").Model(CreateModel{}).Create(&newUser).Error
	return newUser, err
}

func (user Service) Update(id string, newUser UpdateModel) error {
	var err error

	err = database.DB.Table("users").Where("id = ?", id).Updates(&newUser).Error
	return err
}

func (user Service) Erase(id string) error {
	err := database.DB.Table("users").Delete(ViewModel{}, id).Error
	return err
}
