package adminservice

import (
	"errors"
	"go-hrs/config/database"
	"go-hrs/helpers/requesthelper"
	"go-hrs/models/adminmodel"
	"time"
)

func GetAdmins(query string) ([]adminmodel.ViewAdmin, error) {
	var qryMap = requesthelper.ProcessQry(query)
	var admins []adminmodel.ViewAdmin
	var err error = nil

	err = database.DB.Table("admin").Preload("User").Preload("Role").Where(qryMap).Find(&admins).Error
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

func DeleteAdmin(adminID string, deletedBy string) error {
	var whereMap = make(map[string]any)
	whereMap["admin_id"] = adminID
	whereMap["is_deleted"] = "0"

	var deleteMap = make(map[string]any)
	deleteMap["is_deleted"] = "1"
	deleteMap["deleted_at"] = time.Now()
	deleteMap["deleted_by"] = deletedBy

	delete := database.DB.Table("admin").Where(whereMap).Updates(deleteMap)
	if delete.Error != nil {
		return delete.Error
	}

	if delete.RowsAffected == 0 {
		return errors.New("No admin was deleted!")
	}

	return nil
}
