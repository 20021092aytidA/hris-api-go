package applicantdetailcontroller

import (
	"fmt"
	"go-hrs/helpers/jwthelper"
	"go-hrs/models/applicantdetailmodel"
	"go-hrs/services/applicantdetailservice"
	"go-hrs/services/userservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApplicantDetails(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "user")
	if !isValidJWT {
		return
	}

	var applicantDetails []applicantdetailmodel.ViewApplicantDetail
	var err error = nil
	query := c.Request.URL.RawQuery

	applicantDetails, err = applicantdetailservice.GetApplicantDetails(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "Failed to retrieve applicant details!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully retrieve applicant details!",
		"data":    applicantDetails,
	})
}

func CreateApplicantDetail(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "user")
	if !isValidJWT {
		return
	}

	var applicantDetail applicantdetailmodel.CreateApplicantDetail

	if err := c.ShouldBind(&applicantDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to create new applicant detail!",
			"description": "Missing body!",
		})
		return
	}

	//USER
	user, _ := userservice.GetUsers(fmt.Sprintf("user_id=%v", *applicantDetail.UserID))
	if len(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":      http.StatusNotFound,
			"message":     "Failed to create new applicant detail!",
			"description": "User not found!",
		})
		return
	}

	if err := applicantdetailservice.CreateApplicantDetail(applicantDetail); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to create new applicant detail!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully created new applicant detail!",
	})
}

func UpdateApplicantDetail(c *gin.Context) {
	isValidJWT := jwthelper.CheckAndValidateToken(c, "user")
	if !isValidJWT {
		return
	}

	id := c.Param("id")
	var applicantDetail applicantdetailmodel.UpdateApplicantDetail

	if err := c.ShouldBind(&applicantDetail); err != nil || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "Failed to update applicant detail!",
			"description": "Missing body or param!",
		})
	}

	//EXIST
	existAppDet, _ := applicantdetailservice.GetApplicantDetails(fmt.Sprintf("applicant_detail_id=%v", id))
	if len(existAppDet) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":      http.StatusNotFound,
			"message":     "Failed to update applicant detial!",
			"description": "Applicatn detail not found!",
		})
		return
	}

	//USER
	if applicantDetail.UserID != nil {
		user, _ := userservice.GetUsers(fmt.Sprintf("user_id=%v", *applicantDetail.UserID))
		if len(user) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":      http.StatusNotFound,
				"message":     "Failed to update applicant detial!",
				"description": "User not found!",
			})
			return
		}
	}

	if err := applicantdetailservice.UpdateApplicantDetail(id, applicantDetail); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":      http.StatusConflict,
			"message":     "Failed to update applicant detail!",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully updated applicant detail!",
	})
}
