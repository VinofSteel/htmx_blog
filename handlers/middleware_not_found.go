package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/vinofsteel/htmx_blog/views"
)

func (cfg *Config) MiddlewareNotFound(c *fiber.Ctx) error {
	slug := c.OriginalURL()[1:]

	return cfg.render(c, views.NotFound(slug), templ.WithStatus(fiber.StatusNotFound))
}
