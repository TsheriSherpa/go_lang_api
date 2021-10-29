package database

import (
	"fmt"

	"github.com/tsheri/go-fiber/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabaseConnection() error {
	connectionString, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to generate mysql connection string")
	}
	fmt.Println(connectionString)
	DB, err = gorm.Open(mysql.Open(connectionString))
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}
	return nil
}
