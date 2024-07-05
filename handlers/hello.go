package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinofsteel/htmx_blog/templates"
)

func (cfg *Config) RenderHello(c *fiber.Ctx) error {
	name := c.Params("name")
	
	if name == "" {
		name = "World"
	}

	return cfg.render(c, templates.Home(name))
}