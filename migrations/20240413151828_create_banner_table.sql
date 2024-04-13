-- +goose Up
CREATE TABLE banner (
    id SERIAL PRIMARY KEY,
    feature_id INT UNIQUE,
    content JSONB,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (feature_id) REFERENCES feature(id)
);

-- +goose Down
DROP TABLE IF EXISTS banner;
