package main

import (
	"fmt"
	"go-hris/api"
	"go-hris/config/database"
	"go-hris/config/env"
	"go-hris/middleware/cors"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	env.Load()
	database.ConnectMySQL()
	app := gin.Default()
	app.Use(cors.Setup())
	api.InitRoute(app)
	app.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT")))
}
