package banner_dto

import (
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/guregu/null"
)

type BannerResponse struct {
	ID        int       `json:"id"`
	IsActive  bool      `json:"is_active"`
	Content   string    `json:"content"`
	FeatureID int       `json:"feature_id"`
	TagIDs    []int     `json:"tag_ids"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

type ContentResponse struct {
	Content string `json:"content"`
}

func NewContentResponseFromDomain(content string) (*ContentResponse, error) {
	return &ContentResponse{
		Content: content,
	}, nil
}

func NewBannerResponseFromDomain(b banner.Banner) (*BannerResponse, error) {
	return &BannerResponse{
		ID:        b.ID,
		IsActive:  b.IsActive,
		Content:   b.Content,
		FeatureID: b.FeatureID,
		TagIDs:    b.TagIDs,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}, nil
}

func NewBannerResponsesFromDomain(banners []banner.Banner) ([]BannerResponse, error) {
	responses := make([]BannerResponse, len(banners))
	for i, b := range banners {
		responses[i] = BannerResponse{
			ID:        b.ID,
			IsActive:  b.IsActive,
			Content:   b.Content,
			FeatureID: b.FeatureID,
			TagIDs:    b.TagIDs,
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
		}
	}
	return responses, nil
}
