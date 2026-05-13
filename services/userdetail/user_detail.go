package userdetail

import (
	"go-hris/config/database"
	"go-hris/models/userdetail"
	"strconv"
)

func Find(qry map[string]any) ([]userdetail.View, error) {
	var listUserDetail []userdetail.View
	var err error

	err = database.DB.Preload("User").Preload("User.Role").Table("user_details").Model([]userdetail.View{}).Where(qry).Find(&listUserDetail).Error
	return listUserDetail, err
}

func Create(newUserDetail userdetail.Create) error {
	var err error

	err = database.DB.Table("user_details").Model(userdetail.Create{}).Create(newUserDetail).Error
	return err
}

func Update(id string, newUserDetail userdetail.Update) error {
	var err error

	err = database.DB.Table("user_details").Model(userdetail.Update{}).Where("id = ?", id).Updates(newUserDetail).Error
	return err
}

func Erase(id string) error {
	numID, _ := strconv.Atoi(id)
	var err error

	err = database.DB.Table("user_details").Delete(userdetail.Delete{
		Id: numID,
	}).Error
	return err
}
