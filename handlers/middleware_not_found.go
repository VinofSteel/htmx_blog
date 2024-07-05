package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/vinofsteel/htmx_blog/views"
)

func (cfg *Config) Middleware_NotFound(c *fiber.Ctx) error {
	return cfg.render(c, views.NotFound(), templ.WithStatus(fiber.StatusNotFound))
}
