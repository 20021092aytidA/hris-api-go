package role

import (
	"go-hris/config/database"
	"go-hris/models/role"
)

func Find(param role.AllParam) ([]role.View, error) {
	var listRole []role.View
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Table("roles").Model([]role.View{}).Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listRole).Error
	return listRole, err
}

func Create(newRole role.Create) error {
	err := database.DB.Table("roles").Model(role.Create{}).Create(&newRole).Error
	return err
}

func Update(id string, newRole role.Update) error {
	err := database.DB.Table("roles").Model(role.Update{}).Where("id = ?", id).Updates(&newRole).Error
	return err
}

func Erase(id string) error {
	err := database.DB.Table("roles").Delete(role.View{}, id).Error
	return err
}
