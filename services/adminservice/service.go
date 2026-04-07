package adminservice

import (
	"errors"
	"go-hrs/config/database"
	"go-hrs/helpers/requesthelper"
	"go-hrs/models/adminmodel"
)

func GetAdmins(query string) ([]adminmodel.ViewAdmin, error) {
	var qryMap = requesthelper.ProcessQry(query)
	var admins []adminmodel.ViewAdmin
	var err error = nil

	err = database.DB.Table("admin").Where(qryMap).Find(&admins).Error
	if err != nil {
		return admins, err
	}

	return admins, nil
}

func CreateAdmin(admin adminmodel.CreateAdmin) error {
	post := database.DB.Table("admin").Create(admin)
	if post.Error != nil {
		return post.Error
	}

	return nil
}

func UpdateAdmin(adminID string, adminUpdates adminmodel.UpdateAdmin) error {
	var whereMap = make(map[string]any)
	whereMap["admin_id"] = adminID
	whereMap["is_deleted"] = "0"
	put := database.DB.Table("admin").Where(whereMap).Updates(adminUpdates)
	if put.Error != nil {
		return put.Error
	}

	if put.RowsAffected == 0 {
		return errors.New("No admin was updated!")
	}

	return nil
}
