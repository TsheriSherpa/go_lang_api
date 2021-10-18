package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:root@tcp(db:3306)/go_webapp?charset=utf8&parseTime=true"

func InitDatabaseConnection() (*gorm.DB, error) {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}
	return DB, nil
}
