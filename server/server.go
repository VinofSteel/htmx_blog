package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(config fiber.Config) *fiber.App {
	app := fiber.New(config)
	app.Static("/", "./public") // Serving static files from this folder

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "IP+PORT: ${ip}:${port} | METHOD: ${method} | STATUS: ${status} | PATH: ${path}\n",
	}))

	// Router
	routerInitializer(app)
	return app
}