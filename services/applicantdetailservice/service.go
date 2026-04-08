package applicantdetailservice

import (
	"go-hrs/config/database"
	"go-hrs/helpers/requesthelper"
	"go-hrs/models/applicantdetailmodel"
)

func GetApplicantDetails(query string) ([]applicantdetailmodel.ViewApplicantDetail, error) {
	var qryMap = requesthelper.ProcessQry(query)
	var applicantDetails []applicantdetailmodel.ViewApplicantDetail
	var err error = nil

	err = database.DB.Table("applicant_detail").Where(qryMap).Find(&applicantDetails).Error
	if err != nil {
		return applicantDetails, err
	}

	return applicantDetails, nil
}

func CreateApplicantDetail(applicantDetail applicantdetailmodel.CreateApplicantDetail) error {
	create := database.DB.Table("applicant_detail").Create(applicantDetail)

	if create.Error != nil {
		return create.Error
	}

	return nil
}
