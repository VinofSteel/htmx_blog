package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
	"github.com/vinofsteel/templ_blog/internal/database"
)

type Article struct {
	ID        uuid.UUID       `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt sql.NullTime    `json:"deleted_at"`
	Slug      string          `json:"slug"`
	Title     string          `json:"title"`
	Author    string          `json:"author"`
	Content   json.RawMessage `json:"content"`
}

type Embed struct {
	Key   string
	Value interface{}
}

type Op struct {
	Insert      interface{}            `json:"insert,omitempty"`
	InsertEmbed *Embed                 `json:"-"`
	Retain      *int                   `json:"retain,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	Delete      *int                   `json:"delete,omitempty"`
}

func (cfg *Config) ArticlesCreate(c *fiber.Ctx) error {
	c.Accepts("application/json")

	type parameters struct {
		Title          string `json:"title"`
		Author         string `json:"author"`
		Slug           string `json:"slug"`
		ArticleContent []Op   `json:"article_content"`
	}

	params := parameters{}
	if err := c.BodyParser(&params); err != nil {
		log.Println("Error parsing JSON body: ", err)
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "error parsing body in request",
		}
	}

	// errors := cfg.Validator.ValidateData(params)
	// if errors != nil {
	// 	return errors
	// }

	// Check if an article with the same slug already exists
	existingArticle, err := cfg.DB.ListArticleBySlug(c.Context(), params.Slug)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error trying to get an article by slug in ArticlesCreate: ", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "unknown error",
		}
	}

	if existingArticle.ID != uuid.Nil {
		log.Println("Trying to create article with existing slug in DB")
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "article with this slug already exists",
		}
	}

	// Marshal the article content into JSON
	contentJSON, err := json.Marshal(params.ArticleContent)
	if err != nil {
		log.Println("Error marshaling article content to JSON: ", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "error processing article content",
		}
	}

	articleParams := database.CreateArticleParams{
		Slug:    params.Slug,
		Title:   params.Title,
		Author:  params.Author,
		Content: pqtype.NullRawMessage{RawMessage: contentJSON, Valid: true},
	}

	newArticle, err := cfg.DB.CreateArticle(c.Context(), articleParams)
	if err != nil {
		log.Println("Error creating article: ", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "error creating article",
		}
	}

	return c.Status(fiber.StatusCreated).JSON(databaseArticleToHandlerArticle(newArticle))
}

func (cfg *Config) ArticlesListBySlug(c *fiber.Ctx) error {
	c.Accepts("application/json")

	slug := c.Params("slug")
	article, err := cfg.DB.ListArticleBySlug(c.Context(), slug)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Trying to get a non-existing article from db")
			return &fiber.Error{
				Code:    fiber.StatusNotFound,
				Message: "article with this slug not found",
			}
		}

		log.Println("Error trying to get an article by slug in ArticlesCreate: ", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: "unknown error",
		}
	}

	log.Println(databaseArticleToHandlerArticle(article), "ARTICLE")

	return c.Status(fiber.StatusOK).JSON(databaseArticleToHandlerArticle(article))
}

// Utilities
func databaseArticleToHandlerArticle(article database.Article) Article {
	return Article{
		ID:        article.ID,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		Slug:      article.Slug,
		Title:     article.Title,
		Author:    article.Author,
		Content:   article.Content.RawMessage,
	}
}
