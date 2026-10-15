package request

import (
	"go-hris/config/database"
)

type Service struct{}

func (req Service) Find(param AllParam) ([]ViewModel, error) {
	var listRequest []ViewModel
	offset := (param.Pagination.Page - 1) * param.Pagination.Limit
	err := database.DB.Table("requests").Where(&param.AllowedParam).Offset(offset).Limit(param.Pagination.Limit).Find(&listRequest).Error
	return listRequest, err
}

func (req Service) Create(newRequest CreateModel) error {
	err := database.DB.Table("requests").Model(CreateModel{}).Create(&newRequest).Error
	return err
}

func (req Service) Update(id string, updateRequest UpdateModel) error {
	err := database.DB.Table("requests").Model(UpdateModel{}).Where("id = ?", id).Updates(&updateRequest).Error
	return err
}

func (req Service) Erase(id string) error {
	err := database.DB.Table("requests").Delete(ViewModel{}, id).Error
	return err
}
