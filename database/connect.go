package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DB_USERNAME = "tanusree"
const DB_PASSWORD = "password@12345"
const DB_NAME = "asthatrade_all_db"
const DB_HOST = "astha-dev.cwx3kboqzurw.ap-south-1.rds.amazonaws.com"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = GetDB()
	return Db
}

func GetDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}
	return db
}
