package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/vinofsteel/htmx_blog/handlers"
)

func main() {	
	// Initializers
	// Initializing environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initializng our db
	// db := dbInitializer()

	// Initializing the validator
	// validate := validator.New(validator.WithRequiredStructEnabled())

	// validator := validation.New(validate)
	// dbQueries := database.New(db)

	// handlersConfig := handlers.Config{
	// 	DB:        dbQueries,
	// 	Validator: validator,
	// }

	// Handlers config
	handlersConfig := handlers.Config{}

	// New fiber app
	config := fiber.Config{
		AppName:      "HTMX Blog",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorHandler: globalErrorHandler,
	}
	app := fiber.New(config)

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "IP+PORT: ${ip}:${port} | METHOD: ${method} | STATUS: ${status} | PATH: ${path}\n",
	}))

	// Routes
	app.Get("/:name?", handlersConfig.RenderHello)

	// Fallback 404 middleware, catches any route that does not have a handler
	app.Use(handlersConfig.Middleware_NotFound)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func globalErrorHandler(c *fiber.Ctx, err error) error {
	type serverError struct {
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors,omitempty"`
	}

	switch e := err.(type) {
	case *fiber.Error:
		return c.Status(e.Code).JSON(serverError{
			Message: e.Message,
		})
	// case *validation.ValidationError:
	// 	return c.Status(fiber.StatusBadRequest).JSON(serverError{
	// 		Message: e.Error(),
	// 		Errors:  e.Errors,
	// 	})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(serverError{
			Message: err.Error(),
		})
	}
}
