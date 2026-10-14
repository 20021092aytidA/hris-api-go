package main

import (
	"fmt"
	"go-hris/api"
	"go-hris/config/database"
	"go-hris/config/env"
	"go-hris/middleware/cors"
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
	api.InitRoute(app)

	app.Run(fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")))
}
