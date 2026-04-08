package userservice

import (
	"errors"
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

func UpdateUser(userID string, user usermodel.UpdateUser) error {
	var whereMap = make(map[string]any)
	whereMap["user_id"] = userID
	whereMap["is_deleted"] = "0"

	update := database.DB.Table("user").Where(whereMap).Updates(user)
	if update.Error != nil {
		return update.Error
	}

	if update.RowsAffected == 0 {
		return errors.New("No user was updated!")
	}

	return nil
}
