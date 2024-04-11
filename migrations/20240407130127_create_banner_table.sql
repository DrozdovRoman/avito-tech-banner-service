-- +goose Up
CREATE TABLE banner (
    id SERIAL PRIMARY KEY,
    is_active BOOLEAN NOT NULL,
    content JSONB,
    feature_id INT,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    FOREIGN KEY (feature_id) REFERENCES feature(id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS banner;
