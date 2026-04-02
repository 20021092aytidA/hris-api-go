package main

import (
	"fmt"
	"go-hrs/config/database"
	"go-hrs/config/env"
	"go-hrs/routes/adminroute"
	"go-hrs/routes/applicantdetailroute"
	"go-hrs/routes/roleroute"
	"go-hrs/routes/userroute"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := env.LoadENV(); err != nil {
		panic(strings.ToUpper(err.Error()))
	}
	if err := database.LoadMySQL(); err != nil {
		panic(strings.ToUpper(err.Error()))
	}
	app := gin.Default()

	// ROUTES
	roleroute.InitRoute(app)
	adminroute.InitRoute(app)
	applicantdetailroute.InitRoute(app)
	userroute.InitRoute(app)

	app.Run(fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")))
}
