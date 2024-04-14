package banner_dto

import (
	"encoding/json"
	"errors"
)

type PatchBannerRequest struct {
	TagIDs    []int  `json:"tag_ids"`
	FeatureID int    `json:"feature_id"`
	Content   string `json:"content"`
	IsActive  bool   `json:"is_active"`
}

func (c *PatchBannerRequest) Validate() error {
	if c.Content != "" {
		var jsonCheck interface{}
		if err := json.Unmarshal([]byte(c.Content), &jsonCheck); err != nil {
			return errors.New("content must be valid JSON")
		}
	}
	return nil
}
