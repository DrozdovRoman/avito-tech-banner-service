-- +goose Up
CREATE TABLE banner_tag (
    banner_id INT,
    tag_id INT,
    PRIMARY KEY (banner_id, tag_id),
    FOREIGN KEY (banner_id) REFERENCES banner(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS banner_tag;
