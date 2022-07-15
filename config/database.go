package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

const (
	DBUser     = "root"
	DBPassword = ""
	DBName     = "go_basic"
	DBHost     = "127.0.0.1"
	DBPort     = "3306"
)

func connectDB() *gorm.DB {
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
