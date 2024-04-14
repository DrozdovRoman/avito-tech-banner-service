package banner

import (
	"encoding/json"
	"errors"
	"github.com/guregu/null"
	"time"
)

// Errors
var (
	ErrNoTagIDs    = errors.New("at least one tag ID must be provided")
	ErrNoFeatureID = errors.New("feature ID must be provided")
	ErrJSONMarshal = errors.New("error marshalling content into JSON")
)

type Banner struct {
	ID        int       `json:"id"`
	TagIDs    []int     `json:"tag_ids" db:"tag_ids"`
	FeatureID int       `json:"feature_id"`
	Content   string    `json:"content"`
	IsActive  bool      `json:"is_active"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func NewBanner(tagIDs []int, featureID int, content string, isActive bool) (*Banner, error) {
	if len(tagIDs) == 0 {
		return nil, ErrNoTagIDs
	}

	if featureID == 0 {
		return nil, ErrNoFeatureID
	}

	_, err := json.Marshal(content)
	if err != nil {
		return nil, ErrJSONMarshal
	}

	return &Banner{
		TagIDs:    tagIDs,
		FeatureID: featureID,
		Content:   content,
		IsActive:  isActive,
		UpdatedAt: null.TimeFrom(time.Now()),
		CreatedAt: null.TimeFrom(time.Now()),
	}, nil

}

func (b *Banner) GetID() int {
	return b.ID
}

func (b *Banner) GetContent() string {
	return b.Content
}

func (b *Banner) SetContent(content string) {
	b.Content = content
}

func (b *Banner) GetTagIDs() []int {
	return b.TagIDs
}

func (b *Banner) SetTagIDs(tagIDs []int) {
	b.TagIDs = tagIDs
}

func (b *Banner) GetFeature() int {
	return b.FeatureID
}

func (b *Banner) SetFeature(feature int) {
	b.FeatureID = feature
}

func (b *Banner) GetIsActive() bool {
	return b.IsActive
}

func (b *Banner) SetIsActive(isActive bool) {
	b.IsActive = isActive
}

func (b *Banner) SetUpdatedAt(updatedAt time.Time) {
	b.UpdatedAt = null.TimeFrom(updatedAt)
}

func (b *Banner) GetType() string {
	return "banner"
}
