package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tsheri/go-fiber/pkg/utils"
	"github.com/tsheri/go-fiber/platform/database"
	"github.com/tsheri/go-fiber/user/models"
)

// CreateUser func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param firstname body string true "FirstName"
// @Param lastname body string true "LastName"
// @Param email body string true "Email"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /v1/user [post]
func CreateUser(c *fiber.Ctx) error {

	// Create new User struct
	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Set initialized default data for user:
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.Status = 1 // 0 == draft, 1 == active

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": true,
			"msg":    utils.ValidatorErrors(err),
		})
	}

	// Create user by given ID.
	if err := db.CreateUser(user); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": true,
			"msg":    err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"status": true,
		"msg":    "user created",
		"user":   user,
	})
}
