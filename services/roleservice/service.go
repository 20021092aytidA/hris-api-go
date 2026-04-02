package roleservice

import (
	"go-hrs/config/database"
	"go-hrs/helpers/requesthelper"
	"go-hrs/models/rolemodel"
)

func GetRoles(query string) ([]rolemodel.ViewRole, error) {
	var qryMap = requesthelper.ProcessQry(query)
	var roles []rolemodel.ViewRole
	var err error = nil

	err = database.DB.Table("admin").Where(qryMap).Find(&roles).Error
	if err != nil {
		return roles, err
	}

	return roles, nil
}
