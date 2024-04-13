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
	ID        int             `json:"id"`
	TagIDs    []int           `json:"tagIDs"`
	FeatureID null.Int        `json:"featureID"`
	Content   json.RawMessage `json:"content"`
	IsActive  bool            `json:"isActive"`
	CreatedAt null.Time       `json:"createdAt"`
	UpdatedAt null.Time       `json:"updatedAt"`
}

func NewBanner(tagIDs []int, featureID int, content string, isActive bool) (*Banner, error) {
	if len(tagIDs) == 0 {
		return nil, ErrNoTagIDs
	}

	if featureID == 0 {
		return nil, ErrNoFeatureID
	}

	bannerContent, err := json.Marshal(content)
	if err != nil {
		return nil, ErrJSONMarshal
	}

	return &Banner{
		TagIDs:    tagIDs,
		FeatureID: null.NewInt(int64(featureID), true),
		Content:   bannerContent,
		IsActive:  isActive,
		UpdatedAt: null.TimeFrom(time.Now()),
		CreatedAt: null.TimeFrom(time.Now()),
	}, nil

}

func (b *Banner) GetID() int {
	return b.ID
}

func (b *Banner) GetContent() json.RawMessage {
	return b.Content
}

func (b *Banner) SetContent(content json.RawMessage) {
	b.Content = content
}

func (b *Banner) GetTagIDs() []int {
	return b.TagIDs
}

func (b *Banner) AddTagID(tag int) {
	for _, existingTag := range b.TagIDs {
		if existingTag == tag {
			return
		}
	}
	b.TagIDs = append(b.TagIDs, tag)
}

func (b *Banner) RemoveTagID(tag int) {
	for i, existingTag := range b.TagIDs {
		if existingTag == tag {
			b.TagIDs = append(b.TagIDs[:i], b.TagIDs[i+1:]...)
			return
		}
	}
}

func (b *Banner) GetFeatures() null.Int {
	return b.FeatureID
}

func (b *Banner) SetFeatures(features null.Int) {
	b.FeatureID = features
}

func (b *Banner) GetIsActive() bool {
	return b.IsActive
}

func (b *Banner) SetIsActive(isActive bool) {
	b.IsActive = isActive
}

func (b *Banner) GetType() string {
	return "banner"
}
