package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tsheri/go-fiber/pkg/configs"
	"github.com/tsheri/go-fiber/pkg/middleware"
	"github.com/tsheri/go-fiber/pkg/routes"
	"github.com/tsheri/go-fiber/pkg/utils"
)

func index(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.Render("index", fiber.Map{
		"hello": "world",
	})
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config) // Define new fiber app
	app.Get("/", index)
	app.Static("/", "./public")     // set static files location
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)      // Register a route for API Docs (Swagger).
	routes.RegisterApiRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app)     // Register route for 404 Error.

	if utils.GetEnv("STAGE_STATUS", "") == "dev" {
		utils.StartServer(app)
	} else {
		// Start server (with or without graceful shutdown).
		utils.StartServerWithGracefulShutdown(app)
	}
}
