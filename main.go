package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tsheri/go-fiber/config"
	"github.com/tsheri/go-fiber/database"
	"github.com/tsheri/go-fiber/database/migration"
	"github.com/tsheri/go-fiber/user"
)

func hello(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString("<h1>Hello, World</h1>")
}

func setupApiRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/users", user.GetUsers)
	api.Get("/user/:id", user.GetUser)
	api.Post("/user", user.SaveUser)
	api.Delete("/user/:id", user.DeleteUser)
	api.Put("/user/:id", user.UpdateUser)
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	conf := config.New()
	database.InitDatabaseConnection(conf)
	migration.Migrate(database.DB)
}

func main() {
	app := fiber.New()
	setupApiRoutes(app)
	app.Get("/", hello)
	app.Listen(":5000")
}
