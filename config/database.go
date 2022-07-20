package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func connectDB() *gorm.DB {
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASS")
	DBName := os.Getenv("DB_NAME")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")

	DBurl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(DBurl), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}
	return db
}

func InitDB() *gorm.DB {
	DB = connectDB()
	return DB
}
