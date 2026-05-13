package request

import (
	"go-hris/config/database"
	"go-hris/models/request"
)

func Find(qry map[string]any) ([]request.View, error) {
	var listOfRequest []request.View
	var err error
	err = database.DB.Table("requests").Model([]request.View{}).Where(qry).Find(&listOfRequest).Error

	return listOfRequest, err
}

func Create(newRequest request.Create) error {
	var err error
	err = database.DB.Table("requests").Model(request.Create{}).Create(newRequest).Error

	return err
}

func Erase(id string) error {
	var err error
	var deleteMap = make(map[string]string)
	deleteMap["id"] = id

	err = database.DB.Table("requests").Delete(deleteMap).Error

	return err
}
