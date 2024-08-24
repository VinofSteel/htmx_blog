-- +goose Up
ALTER TABLE articles ADD COLUMN title TEXT NOT NULL;

-- +goose Down
ALTER TABLE articles DROP COLUMN title;
