package banner_dto

import (
	"encoding/json"
	"errors"
)

type CreateBannerRequest struct {
	TagIDs    []int  `json:"tag_ids"`
	FeatureID int    `json:"feature_id"`
	Content   string `json:"content"`
	IsActive  bool   `json:"is_active"`
}

func (c *CreateBannerRequest) Validate() error {
	if c.FeatureID == 0 {
		return errors.New("feature ID is required")
	}
	if len(c.TagIDs) == 0 {
		return errors.New("at least one tag ID is required")
	}
	if c.Content == "" {
		return errors.New("content cannot be empty")
	}

	var jsonCheck interface{}
	if err := json.Unmarshal([]byte(c.Content), &jsonCheck); err != nil {
		return errors.New("content must be valid JSON")
	}

	return nil
}
