package banner

import (
	"context"
	"encoding/json"
)

type Repository interface {
	GetBanner(ctx context.Context, id int) (*Banner, error)
	GetBanners(ctx context.Context, tagID, featureID, limit, offset int) ([]Banner, error)
	GetActiveBannerContentByTagAndFeature(ctx context.Context, tagID int, featureID int) (json.RawMessage, error)
	AddBanner(ctx context.Context, banner *Banner) (int, error)
	UpdateBanner(ctx context.Context, banner *Banner) error
	DeleteBanner(ctx context.Context, id int) error
	BannerExists(ctx context.Context, id int) (bool, error)
}
