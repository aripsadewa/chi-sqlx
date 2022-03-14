package config

import (
	"fmt"
	"rest_api/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dB *gorm.DB

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	dbUser := utils.EnvConfigs.DbUser
	dbPass := utils.EnvConfigs.DbPass
	dbHost := utils.EnvConfigs.DbHost
	dbName := utils.EnvConfigs.DbName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	sqlDb, err := db.DB()
	sqlDb.SetMaxOpenConns(100)
	dB = db
	return dB
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
