package banner

import (
	"context"
	"encoding/json"
)

type Repository interface {
	GetActiveBannerContentByTagAndFeature(ctx context.Context, tagID int, featureID int) (json.RawMessage, error)
	GetAll() ([]Banner, error)
	GetByID(ctx context.Context, bannerID int) (Banner, error)
	Add(banner Banner) error
	Update(banner Banner) error
	Delete(id int) error
}
