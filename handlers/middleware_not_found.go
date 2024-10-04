package handlers

import (
	"database/sql"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vinofsteel/templ_blog/views"
)

func (cfg *Config) MiddlewareNotFound(c *fiber.Ctx) error {
	urlSlug := c.OriginalURL()[1:]
	sanitizedSlug := sanitizeSlug(urlSlug)

	if sanitizedSlug != urlSlug {
		if sanitizedSlug == "" {
			return c.Redirect("/", fiber.StatusMovedPermanently)
		}

		return c.Redirect("/"+sanitizedSlug, fiber.StatusMovedPermanently)
	}

	existingArticle, err := cfg.DB.ListArticleBySlug(c.Context(), sanitizedSlug)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error trying to get an article by slug in MiddlewareNotFound: ", err)
		return cfg.Render(c, views.NotFound(sanitizedSlug), templ.WithStatus(fiber.StatusNotFound))
	}

	if existingArticle.ID == uuid.Nil {
		log.Printf("Article with slug %s not found, rendering not found view...\n", sanitizedSlug)
		return cfg.Render(c, views.NotFound(sanitizedSlug), templ.WithStatus(fiber.StatusNotFound))
	}

	return cfg.Render(c, views.ExistingArticle(existingArticle), templ.WithStatus(fiber.StatusNotFound))
}

func sanitizeSlug(slug string) string {
	// URL-decoding the slug to remove special characters like whitespace
	decodedSlug, err := url.QueryUnescape(slug)
	if err != nil {
		log.Println("Error decoding URL:", err)
		decodedSlug = slug
	}

	decodedSlug = strings.ToLower(decodedSlug)
	decodedSlug = strings.ReplaceAll(decodedSlug, " ", "-")

	// Remove all characters except lowercase letters, numbers, and hyphens
	reg := regexp.MustCompile(`[^a-z0-9-]+`)
	sanitized := reg.ReplaceAllString(decodedSlug, "")

	// Replace multiple consecutive hyphens with a single hyphen
	reg = regexp.MustCompile(`-+`)
	sanitized = reg.ReplaceAllString(sanitized, "-")

	// Trim hyphens from start and end
	sanitized = strings.Trim(sanitized, "-")

	return sanitized
}
