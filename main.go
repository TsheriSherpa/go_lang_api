package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsheri/go-fiber/database"
	"github.com/tsheri/go-fiber/database/migration"
	"github.com/tsheri/go-fiber/user"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)
}

func main() {
	database.InitDatabaseConnection()
	migration.Migrate(database.DB)
	app := fiber.New()
	setupRoutes(app)
	app.Get("/", hello)
	app.Listen(":5000")
}
