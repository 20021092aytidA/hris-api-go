package userservice

import (
	"go-hrs/config/database"
	"go-hrs/helpers/requesthelper"
	"go-hrs/models/usermodel"
)

func GetUsers(query string) ([]usermodel.ViewUser, error) {
	var qryMap = requesthelper.ProcessQry(query)
	var users []usermodel.ViewUser
	var err error = nil

	err = database.DB.Table("user").Where(qryMap).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func CreateUser(user usermodel.CreateUser) error {
	create := database.DB.Table("user").Create(user)

	if create.Error != nil {
		return create.Error
	}

	return nil
}
