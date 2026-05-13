package user

import (
	"go-hris/config/database"
	"go-hris/models/user"
	"strconv"
)

func FindForPassword(qry map[string]any) (user.ViewWithPass, error) {
	var uSer user.ViewWithPass
	var err error

	err = database.DB.Table("users").Model(user.ViewWithPass{}).Where(qry).First(&uSer).Error

	return uSer, err
}

func Find(qry map[string]any) ([]user.View, error) {
	var listUsers []user.View
	var err error

	err = database.DB.Preload("Role").Table("users").Model([]user.View{}).Where(qry).Find(&listUsers).Error
	return listUsers, err
}

func Create(newUser user.Create) error {
	var err error

	err = database.DB.Table("users").Model(user.Create{}).Create(newUser).Error
	return err
}

func Update(id string, newUser user.Update) error {
	var err error

	err = database.DB.Table("users").Model(user.Update{}).Where("id = ?", id).Updates(newUser).Error
	return err
}

func Erase(id string) error {
	numID, _ := strconv.Atoi(id)
	var err error

	err = database.DB.Table("users").Delete(user.Delete{
		Id: numID,
	}).Error

	return err
}
