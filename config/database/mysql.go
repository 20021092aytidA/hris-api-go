package database

import (
	"fmt"
	"go-hris/config/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var gormConfig gorm.Config

func ConnectMySQL() error {
	var dbErr error = nil
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.ENV.DBUser, env.ENV.DBPass, env.ENV.DBHost, env.ENV.DBPort, env.ENV.DBName)
	DB, dbErr = gorm.Open(mysql.Open(dsn), &gormConfig)
	if dbErr != nil {
		return dbErr
	}

	return nil
}
