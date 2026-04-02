package roleservice

import (
	"errors"
	"go-hrs/config/database"
	"go-hrs/helpers/requesthelper"
	"go-hrs/models/rolemodel"
	"time"
)

func GetRoles(query string) ([]rolemodel.ViewRole, error) {
	var qryMap = requesthelper.ProcessQry(query)
	var roles []rolemodel.ViewRole
	var err error = nil

	err = database.DB.Table("role").Where(qryMap).Find(&roles).Error
	if err != nil {
		return roles, err
	}

	return roles, nil
}

func CreateRole(role rolemodel.CreateRole) error {
	post := database.DB.Table("role").Create(role)
	if post.Error != nil {
		return post.Error
	}

	return nil
}

func UpdateRole(roleID string, role rolemodel.UpdateRole) error {
	var whereMap = make(map[string]string)
	whereMap["role_id"] = roleID
	whereMap["is_deleted"] = "0"

	put := database.DB.Table("role").Where(whereMap).Updates(role)
	if put.Error != nil {
		return put.Error
	}

	if put.RowsAffected == 0 {
		return errors.New("No role was updated!")
	}

	return nil
}

func DeleteRole(roleID string, deletedBy string) error {
	var whereMap = make(map[string]string)
	whereMap["role_id"] = roleID
	whereMap["is_deleted"] = "0"

	var deleteMap = make(map[string]any)
	deleteMap["deleted_by"] = deletedBy
	deleteMap["deleted_at"] = time.Now()
	deleteMap["is_deleted"] = "1"

	delete := database.DB.Table("role").Where(whereMap).Updates(deleteMap)
	if delete.Error != nil {
		return delete.Error
	}

	if delete.RowsAffected == 0 {
		return errors.New("No role was deleted!")
	}

	return nil
}
