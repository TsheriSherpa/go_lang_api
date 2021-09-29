package migration

import (
	"github.com/tsheri/go-fiber/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}); err != nil {
		panic("Migration Failed")
	}
	return nil
}
