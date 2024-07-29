package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/vinofsteel/htmx_blog/internal/database"
)

type Config struct {
	DB *database.Queries
	// Validator *validation.Validator
}

// func getAuthHeader(authHeader string) (TokenString, error) {
// 	if authHeader == "" {
// 		return "", errors.New("missing Authorization header")
// 	}

// 	splitAuth := strings.Split(authHeader, " ")
// 	if len(splitAuth) < 2 || splitAuth[0] != "Bearer" {
// 		return "", errors.New("malformed Authorization header")
// 	}

// 	return TokenString(splitAuth[1]), nil
// }

func (cfg *Config) render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

func New(queries *database.Queries) Config {
	return Config{
		DB: queries,
	}
}
