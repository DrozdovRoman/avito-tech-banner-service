-- +goose Up
CREATE TABLE banner (
    id SERIAL PRIMARY KEY,
    is_active BOOLEAN NOT NULL DEFAULT true,
    feature_id INTEGER REFERENCES feature(id) ON DELETE CASCADE NOT NULL,
    content JSONB,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

-- +goose Down
DROP TABLE IF EXISTS banner;
