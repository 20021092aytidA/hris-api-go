package applicantdetailservice

import (
	"errors"
	"go-hrs/config/database"
	"go-hrs/helpers/requesthelper"
	"go-hrs/models/applicantdetailmodel"
	"time"
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

func UpdateApplicantDetail(applicantDetailID string, applicantDetail applicantdetailmodel.UpdateApplicantDetail) error {
	var whereMap = make(map[string]any)
	whereMap["applicant_detail_id"] = applicantDetailID
	whereMap["is_deleted"] = "0"

	update := database.DB.Table("applicant_detail").Where(whereMap).Updates(applicantDetail)
	if update.Error != nil {
		return update.Error
	}

	if update.RowsAffected == 0 {
		return errors.New("No applicant detail was updated!")
	}

	return nil
}

func DeleteApplicantDetail(applicantDetailID string, deletedBy string) error {
	var whereMap = make(map[string]any)
	whereMap["applicant_detail_id"] = applicantDetailID
	whereMap["is_deleted"] = "0"

	var deleteMap = make(map[string]any)
	deleteMap["is_deleted"] = "1"
	deleteMap["deleted_at"] = time.Now()
	deleteMap["deleted_by"] = deletedBy

	delete := database.DB.Table("applicant_detail").Where(whereMap).Updates(deleteMap)
	if delete.Error != nil {
		return delete.Error
	}

	if delete.RowsAffected == 0 {
		return errors.New("No applicant detail was deleted!")
	}

	return nil
}
