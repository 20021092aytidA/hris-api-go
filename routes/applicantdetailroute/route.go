package applicantdetailroute

import (
	"go-hrs/controllers/applicantdetailcontroller"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	c.GET("/hrs-api/applicant-details", applicantdetailcontroller.GetApplicantDetails)
	c.POST("/hrs-api/applicant-detail", applicantdetailcontroller.CreateApplicantDetail)
	c.PUT("/hrs-api/applicant-detail/:id", applicantdetailcontroller.UpdateApplicantDetail)
	c.DELETE("/hrs-api/applicant-detail/:id", applicantdetailcontroller.DeleteApplicantDetail)
}
