package main

import (
	"fmt"
	"go-hrs/config/database"
	"go-hrs/config/env"
	"go-hrs/middleware/cors"
	"go-hrs/routes/request"
	"go-hrs/routes/role"
	"go-hrs/routes/user"
	"go-hrs/routes/userdetail"
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
