package server

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vinofsteel/templ_blog/handlers"
	"github.com/vinofsteel/templ_blog/internal/database"
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

	// Initializing handlers
	handler := handlers.New(queries)

	// API Routes
	app.Get("/api/articles/:slug", handler.ArticlesListBySlug)

	app.Post("/api/articles", handler.ArticlesCreate)

	// Fallback 404 middleware, catches any route that does not have a handler
	app.Use(handler.MiddlewareNotFound)
}
