package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsheri/go-fiber/app/controllers/api"
)

func RegisterApiRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/user", api.CreateUser)
	// route.Get("/user/:id", api.GetUser)
	// route.Post("/user", api.SaveUser)
	// route.Delete("/user/:id", api.DeleteUser)
	// route.Put("/user/:id", api.UpdateUser)
}
