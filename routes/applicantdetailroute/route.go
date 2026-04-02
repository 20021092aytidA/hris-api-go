package applicantdetailroute

import (
	"go-hrs/controllers/applicantdetailcontroller"

	"github.com/gin-gonic/gin"
)

func InitRoute(c *gin.Engine) {
	c.GET("/hrs-api/applicant-details", applicantdetailcontroller.GetApplicantDetails)
}
