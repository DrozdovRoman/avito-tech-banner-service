-- +goose Up
CREATE TABLE feature (
    id SERIAL PRIMARY KEY
);

-- +goose Down
DROP TABLE IF EXISTS feature;
