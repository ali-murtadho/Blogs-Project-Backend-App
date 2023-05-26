package config

import (
	"fmt"
	"project_blog_gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB{
	username := "root"
	password := ""
	host := "tcp(127.0.0.1:3306)"
	database := "blog_gin"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	panic(err.Error())
	}
	db.AutoMigrate(&models.Blog{}, &models.Category{}, &models.User{})
	return db
   
}