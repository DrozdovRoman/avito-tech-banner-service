-- +goose Up
CREATE TABLE tag (
    id SERIAL PRIMARY KEY
);

-- +goose Down
DROP TABLE IF EXISTS tag;
