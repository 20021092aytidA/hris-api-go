package request

import (
	"go-hris/config/database"
)

type Service struct{}

func (req Service) Find(param AllParam) ([]View, error) {
	var listRequest []View
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Table("requests").Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listRequest).Error
	return listRequest, err
}

func (req Service) Create(newRequest Create) error {
	err := database.DB.Table("requests").Model(Create{}).Create(&newRequest).Error
	return err
}

func (req Service) Update(id string, updateRequest Update) error {
	err := database.DB.Table("requests").Model(Update{}).Where("id = ?", id).Updates(&updateRequest).Error
	return err
}

func (req Service) Erase(id string) error {
	err := database.DB.Table("requests").Delete(View{}, id).Error
	return err
}
