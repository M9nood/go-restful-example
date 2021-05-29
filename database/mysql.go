package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

type DBConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func InitDb(config DBConfig) *gorm.DB {
	Db = connectDB(config)
	return Db
}

func connectDB(config DBConfig) *gorm.DB {
	var err error
	dsn := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@tcp" + "(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	fmt.Println("db connected")
	return db
}

func GetDB() *gorm.DB {
	return Db
}

// func CloseDB() {
// 	Db.Close()
// }
