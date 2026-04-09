package userservice

import (
	"errors"
	"go-hrs/config/database"
	"go-hrs/helpers/requesthelper"
	"go-hrs/models/usermodel"
	"time"
)

func GetUsers(query string) ([]usermodel.ViewUser, error) {
	var qryMap = requesthelper.ProcessQry(query)
	var users []usermodel.ViewUser
	var err error = nil

	err = database.DB.Table("user").Preload("Role").Preload("ApplicantDetail").Where(qryMap).Find(&users).Error
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

func DeleteUser(userID string, deletedBy string) error {
	var whereMap = make(map[string]any)
	whereMap["is_deleted"] = "0"
	whereMap["user_id"] = userID
	var deleteMap = make(map[string]any)
	deleteMap["deleted_by"] = deletedBy
	deleteMap["deleted_at"] = time.Now()
	deleteMap["is_deleted"] = "1"

	delete := database.DB.Table("user").Where(whereMap).Updates(deleteMap)
	if delete.Error != nil {
		return delete.Error
	}

	if delete.RowsAffected == 0 {
		return errors.New("No user was deleted!")
	}

	return nil
}
