package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/vinofsteel/htmx_blog/templates"
)

func (cfg *Config) Middleware_NotFound(c *fiber.Ctx) error {
	return cfg.render(c, templates.NotFound(), templ.WithStatus(fiber.StatusNotFound))
}
