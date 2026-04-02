package roleservice

import (
	"errors"
	"go-hrs/config/database"
	"go-hrs/helpers/requesthelper"
	"go-hrs/models/rolemodel"
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

func UpdateRole(adminID string, role rolemodel.UpdateRole) error {
	var whereMap = make(map[string]string)
	whereMap["role_id"] = adminID
	whereMap["is_deleted"] = "0"

	post := database.DB.Table("role").Where(whereMap).Updates(role)
	if post.Error != nil {
		return post.Error
	}

	if post.RowsAffected == 0 {
		return errors.New("No admin was found!")
	}

	return nil
}
