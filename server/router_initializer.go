package server

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vinofsteel/htmx_blog/handlers"
	"github.com/vinofsteel/htmx_blog/internal/database"
)

func routerInitializer(app *fiber.App) {
	// Initializing environment
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Initializing db
	db := dbInitializer()
	queries := database.New(db)

	handlersConfig := handlers.Config{
		DB: queries,
	}

	// Routes
	app.Get("/:name?", handlersConfig.RenderHello)

	// Fallback 404 middleware, catches any route that does not have a handler
	app.Use(handlersConfig.MiddlewareNotFound)
}
