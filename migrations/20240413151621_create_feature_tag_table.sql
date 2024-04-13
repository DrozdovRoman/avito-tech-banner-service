-- +goose Up
CREATE TABLE feature_tag (
    tag_id INT,
    feature_id INT,
    PRIMARY KEY (tag_id, feature_id),
    FOREIGN KEY (tag_id) REFERENCES tag(id),
    FOREIGN KEY (feature_id) REFERENCES feature(id)
);

-- +goose Down
DROP TABLE IF EXISTS feature_tag;
