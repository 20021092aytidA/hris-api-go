package role

import (
	"go-hris/config/database"
	"go-hris/models/role"
	"strconv"
)

func Find(qry map[string]any) ([]role.View, error) {
	var listRole []role.View
	var err error

	err = database.DB.Table("roles").Model([]role.View{}).Where(qry).Find(&listRole).Error

	return listRole, err
}

func Create(newRole role.Create) error {
	var err error
	err = database.DB.Table("roles").Model(role.Create{}).Create(newRole).Error

	return err
}

func Update(id string, newRole role.Update) error {
	var err error
	err = database.DB.Table("roles").Model(role.Update{}).Where("id = ?", id).Updates(newRole).Error

	return err
}

func Erase(id string) error {
	numID, _ := strconv.Atoi(id)
	var err error
	err = database.DB.Table("roles").Delete(role.Delete{
		Id: numID,
	}).Error

	return err
}
