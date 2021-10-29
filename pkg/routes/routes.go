package routes

import (
	"github.com/gofiber/fiber/v2"
	user "github.com/tsheri/go-fiber/user/routes"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

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
