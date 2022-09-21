package routes

import (
	"github.com/gofiber/fiber/v2"
	user "github.com/tsheri/go-fiber/app/routes"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

func index(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.Render("index", fiber.Map{
		"hello": "world",
	})
}

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/swagger")

	// Routes for GET method:
	route.Get("*", swagger.Handler) // get one user by ID
}

func RegisterApiRoutes(a *fiber.App) {
	user.RegisterApiRoutes(a)
}

func RegisterWebRoutes(a *fiber.App) {
	route := a.Group("/admin")
	route.Get("", index)
}

func RegisterRoutes(a *fiber.App) {
	RegisterWebRoutes(a)
	SwaggerRoute(a)
	RegisterApiRoutes(a)
	NotFoundRoute(a)
}
