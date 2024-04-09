package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	table     = "banner"
	id        = "id"
	isActive  = "is_active"
	content   = "content"
	featureID = "feature_id"
	createdAt = "created_at"
	updatedAt = "updated_at"
)

type BannerRepository struct {
	db db.Client
}

func NewBannerRepository(client db.Client) banner.Repository {
	return &BannerRepository{db: client}
}

func (b *BannerRepository) GetAll() ([]banner.Banner, error) {
	return nil, nil
}

func (b *BannerRepository) GetByID(ctx context.Context, bannerID int) (banner.Banner, error) {
	builderSelectByID := sq.Select(id, isActive, content, featureID, createdAt, updatedAt).
		From(table).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{id: bannerID}).
		Limit(1)

	query, args, err := builderSelectByID.ToSql()
	if err != nil {
		return banner.Banner{}, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "banner_repository.Get",
		QueryRaw: query,
	}
	var result banner.Banner
	err = b.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&result.ID,
		&result.Tag,
		&result.Features,
		&result.Content,
		&result.IsActive,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return banner.Banner{}, fmt.Errorf("banner with ID %d not found", bannerID)
		}
		return banner.Banner{}, fmt.Errorf("failed to scan row: %v", err)
	}

	return result, nil
}

func (b *BannerRepository) Add(banner banner.Banner) error {
	_ = banner
	return nil
}

func (b *BannerRepository) Update(banner banner.Banner) error {
	_ = banner
	return nil
}

func (b *BannerRepository) Delete(id int) error {
	_ = id
	return nil
}

func (b *BannerRepository) GetByFeatureID(featureID int) ([]banner.Banner, error) {
	_ = featureID
	return nil, nil
}

func (b *BannerRepository) GetActiveByFeatureID(featureID int) ([]banner.Banner, error) {
	_ = featureID
	return nil, nil
}

func (b *BannerRepository) GetActive() ([]banner.Banner, error) {
	return nil, nil
}
