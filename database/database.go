package database

import (
	"fmt"

	"github.com/tsheri/go-fiber/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabaseConnection(config *config.Config) (*gorm.DB, error) {
	DB, err = gorm.Open(mysql.Open(GetConnectionString(config)), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}
	return DB, nil
}

func GetConnectionString(config *config.Config) string {
	return config.DB.Username + ":" + config.DB.Password + "@" + config.DB.Connection + "(" + config.DB.Host + ":" + config.DB.Port + ")" + "/" + config.DB.Db + "?charset=utf8&parseTime=true"
}
