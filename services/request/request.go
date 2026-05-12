package request

import (
	"go-hrs/config/database"
	"go-hrs/models/request"
)

func Find(qry map[string]any) ([]request.View, error) {
	var listOfRequest []request.View
	var err error = nil
	err = database.DB.Table("requests").Model([]request.View{}).Where(qry).Find(&listOfRequest).Error

	return listOfRequest, err
}
