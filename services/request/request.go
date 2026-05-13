package request

import (
	"go-hris/config/database"
	"go-hris/models/request"
	"strconv"
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
	numID, _ := strconv.Atoi(id)
	var err error

	err = database.DB.Table("requests").Delete(request.Delete{
		Id: numID,
	}).Error

	return err
}
