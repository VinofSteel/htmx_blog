package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vinofsteel/htmx_blog/server"
)

func main() {
	// New fiber app
	config := fiber.Config{
		AppName:      "HTMX Blog",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorHandler: globalErrorHandler,
	}
	
	server := server.New(config)
	log.Fatal(server.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
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
