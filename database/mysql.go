package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/andikanugraha11/go-gin-gorm-restapi/models"
)

func InitMysqlDB() *gorm.DB {
	dsn := "root:hacktiv@tcp(127.0.0.1:3306)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}

	db.AutoMigrate(models.Person{})

	return db
}