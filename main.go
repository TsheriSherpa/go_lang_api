package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tsheri/go-fiber/database"
	"github.com/tsheri/go-fiber/database/migration"
	"github.com/tsheri/go-fiber/user"
)

func hello(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString("<h1>Hello, World d;):sdfsd:: </h1>")
}

func setupApiRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/users", user.GetUsers)
	api.Get("/user/:id", user.GetUser)
	api.Post("/user", user.SaveUser)
	api.Delete("/user/:id", user.DeleteUser)
	api.Put("/user/:id", user.UpdateUser)
}

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("APP_NAME"))
	database.InitDatabaseConnection()
	migration.Migrate(database.DB)
	app := fiber.New()
	setupApiRoutes(app)
	app.Get("/", hello)
	app.Listen(":5000")
}
