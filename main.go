package main

import (
	"fmt"
	"go-hris/config/database"
	"go-hris/config/env"
	"go-hris/middleware/cors"
	"go-hris/routes/request"
	"go-hris/routes/role"
	"go-hris/routes/user"
	"go-hris/routes/userdetail"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := env.Load(); err != nil {
		panic(strings.ToUpper(err.Error()))
	}
	if err := database.ConnectMySQL(); err != nil {
		panic(strings.ToUpper(err.Error()))
	}
	app := gin.Default()
	app.Use(cors.Setup())

	// ROUTES
	request.InitRoute(app)
	role.InitRoute(app)
	user.InitRoute(app)
	userdetail.InitRoute(app)

	app.Run(fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")))
}
