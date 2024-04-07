package banner

import (
	"encoding/json"
	"github.com/guregu/null"
)

type Banner struct {
	ID       int             `json:"id"`
	Tag      []int           `json:"tag"`
	Features null.Int        `json:"features"`
	Content  json.RawMessage `json:"content"`
	IsActive bool            `json:"isActive"`
}

func NewBanner() (*Banner, error) {
	banner := &Banner{}

	return banner, nil
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

func (b *Banner) GetTag() []int {
	return b.Tag
}

func (b *Banner) AddTag(tag int) {
	for _, existingTag := range b.Tag {
		if existingTag == tag {
			return
		}
	}
	b.Tag = append(b.Tag, tag)
}

func (b *Banner) RemoveTag(tag int) {
	for i, existingTag := range b.Tag {
		if existingTag == tag {
			b.Tag = append(b.Tag[:i], b.Tag[i+1:]...)
			return
		}
	}
}

func (b *Banner) GetFeatures() null.Int {
	return b.Features
}

func (b *Banner) SetFeatures(features null.Int) {
	b.Features = features
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
