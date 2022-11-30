package database

import (
	"fmt"

	"github.com/ssss-tantalum/fiber-auth-jwt-sample/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	dsn := "root:root@tcp(127.0.0.1:3306)/fiber_auth_jwt?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{}, &model.Product{})
	fmt.Println("Database Migrated")
}