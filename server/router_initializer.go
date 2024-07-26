package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vinofsteel/htmx_blog/handlers"
)

func routerInitializer(app *fiber.App) {
	// Initializing environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initializing db
	handlersConfig := handlers.Config{}

	// Routes
	app.Get("/:name?", handlersConfig.RenderHello)

	// Fallback 404 middleware, catches any route that does not have a handler
	app.Use(handlersConfig.MiddlewareNotFound)
}
