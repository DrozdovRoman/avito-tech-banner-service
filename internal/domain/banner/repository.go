package banner

import (
	"context"
	"encoding/json"
)

type Repository interface {
	GetActiveBannerContentByTagAndFeature(ctx context.Context, tagID int, featureID int) (json.RawMessage, error)
	Add(ctx context.Context, banner *Banner) (int, error)
}
