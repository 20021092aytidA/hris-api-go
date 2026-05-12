package main

import (
	"fmt"
	"go-hrs/config/database"
	"go-hrs/config/env"
	"go-hrs/middleware/cors"
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

	app.Run(fmt.Sprintf("localhost:%s", os.Getenv("API_PORT")))
}
