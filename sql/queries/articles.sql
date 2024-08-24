-- name: ListAllArticles :many
SELECT * FROM articles ORDER BY $1 OFFSET $2 LIMIT $3;

-- name: ListArticleBySlug :one
SELECT * FROM articles WHERE slug LIKE $1;

-- name: CreateArticle :one
INSERT INTO articles (slug, title, author, content) VALUES ($1, $2, $3, $4) RETURNING *;

