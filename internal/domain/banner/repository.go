package banner

import (
	"context"
	"encoding/json"
)

type Repository interface {
	GetActiveBannerContentByTagAndFeature(ctx context.Context, tagID int, featureID int) (json.RawMessage, error)
	Add(banner *Banner) (int, error)
	GetAll() ([]Banner, error)
	GetByID(ctx context.Context, bannerID int) (Banner, error)
	Update(banner Banner) error
	Delete(id int) error
}
