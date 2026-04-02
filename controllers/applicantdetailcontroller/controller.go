package applicantdetailcontroller

import (
	"go-hrs/models/applicantdetailmodel"
	"go-hrs/services/applicantdetailservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApplicantDetails(c *gin.Context) {
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
