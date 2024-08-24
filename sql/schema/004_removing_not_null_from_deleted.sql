-- +goose Up
ALTER TABLE articles ALTER COLUMN deleted_at DROP NOT NULL;
ALTER TABLE users ALTER COLUMN deleted_at DROP NOT NULL;

-- +goose Down
ALTER TABLE articles ALTER COLUMN deleted_at SET NOT NULL;
ALTER TABLE users ALTER COLUMN deleted_at SET NOT NULL;