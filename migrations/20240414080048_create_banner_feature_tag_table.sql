-- +goose Up
CREATE TABLE banner_feature_tag (
    banner_id INTEGER REFERENCES banner(id) ON DELETE CASCADE NOT NULL,
    tag_id INTEGER REFERENCES tag(id) ON DELETE CASCADE NOT NULL,
    feature_id INTEGER REFERENCES feature(id) ON DELETE CASCADE NOT NULL,
    UNIQUE (tag_id, feature_id)
);

-- +goose Down
DROP TABLE IF EXISTS banner_feature_tag;
