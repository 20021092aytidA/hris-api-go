package request

import (
	"go-hris/config/database"
	"go-hris/models/request"
)

func Find(param request.AllParam) ([]request.View, error) {
	var listRequest []request.View
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Table("requests").Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listRequest).Error
	return listRequest, err
}

func Create(newRequest request.Create) error {
	err := database.DB.Table("requests").Model(request.Create{}).Create(&newRequest).Error
	return err
}

func Update(id string, updateRequest request.Update) error {
	err := database.DB.Table("requests").Model(request.Update{}).Where("id = ?", id).Updates(&updateRequest).Error
	return err
}

func Erase(id string) error {
	err := database.DB.Table("requests").Delete(request.View{}, id).Error
	return err
}
