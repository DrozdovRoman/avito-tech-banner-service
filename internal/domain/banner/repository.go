package banner

import (
	"context"
	"encoding/json"
)

type Repository interface {
	GetBanners(ctx context.Context, tagID, featureID, limit, offset int) ([]Banner, error)
	GetActiveBannerContentByTagAndFeature(ctx context.Context, tagID int, featureID int) (json.RawMessage, error)
	AddBanner(ctx context.Context, banner *Banner) (int, error)
}
