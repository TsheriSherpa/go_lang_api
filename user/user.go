package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsheri/go-fiber/pkg/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `db:"email" json:"email"`
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	database.DB.Find(&user, id)

	if user.Email == "" {
		return c.Status(404).SendString("User Not Found")
	}
	return c.JSON(&user)
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	database.DB.Find(&users)
	return c.JSON(&users)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Create(&user)
	return c.JSON(&user)
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(User)
	id := c.Params("id")

	database.DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User Not Found")
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&user)
	return c.JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	database.DB.First(&user, id)

	if user.Email == "" {
		return c.Status(500).SendString("User Not Found")
	}

	database.DB.Delete(&user)
	return c.SendString("User Deleted Successfully!!!")
}
