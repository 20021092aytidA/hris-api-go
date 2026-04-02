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
