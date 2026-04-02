package adminservice

import (
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
