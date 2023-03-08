package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = GetDB()
	return Db
}

func GetDB() *gorm.DB {
	var err error
	err = godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	fmt.Println("dbUser", dbUser)
	fmt.Println("dbPass", dbPass)
	fmt.Println("dbHost", dbHost)
	fmt.Println("dbName", dbName)
	fmt.Println("dbPort", dbPort)

	// dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable", dbUser, url.QueryEscape(dbPass), dbHost, dbPort, dbName)

	dsn := dbUser + ":" + dbPass + "@tcp" + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}, &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			LogLevel: logger.Silent,
			Colorful: true,
		})})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}
	return db
}
