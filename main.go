package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/redirect/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/tsheri/go-fiber/pkg/configs"
	"github.com/tsheri/go-fiber/pkg/middleware"
	"github.com/tsheri/go-fiber/pkg/routes"
	"github.com/tsheri/go-fiber/pkg/utils"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	url, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		panic("Unable to generate db connection string")
	}

	db, _ := sql.Open("mysql", url)
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		fmt.Println("error, not sent ping to database, %w", err)
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, er := migrate.NewWithDatabaseInstance(
		"file://platform/migrations/",
		"go_webapp",
		driver,
	)
	if er != nil {
		fmt.Println(er.Error())
	}
	m.Steps(2)
}

func main() {

	config := configs.FiberConfig()
	app := fiber.New(config)    // Define new fiber app
	app.Static("/", "./public") // set static files locatigo on

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/admin",
		},
		StatusCode: 301,
	}))

	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.RegisterRoutes(app) // Register a route for API Docs (Swagger).

	if utils.GetEnv("APP_ENV", "") == "production" {
		utils.StartServer(app)
	} else {
		// Start server (with or without graceful shutdown).
		utils.StartServerWithGracefulShutdown(app)
	}
}
