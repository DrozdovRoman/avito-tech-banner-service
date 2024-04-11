package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/db"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

const (
	tableBanner  = "banner"
	colId        = "id"
	colIsActive  = "is_active"
	colContent   = "content"
	colFeatureID = "feature_id"
	colCreatedAt = "created_at"
	colUpdatedAt = "updated_at"

	tableBannerTag = "banner_tag"
	colBannerID    = "banner_id"
	colTagID       = "tag_id"
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
	builderSelectByID := sq.Select(
		colId, colIsActive, colContent, colFeatureID, colCreatedAt, colUpdatedAt, fmt.Sprintf(
			"COALESCE(array_agg(bt.%s) FILTER (WHERE bt.%s IS NOT NULL), '{}') AS tag_ids", colTagID, colTagID),
	).
		From(tableBanner).
		PlaceholderFormat(sq.Dollar).
		LeftJoin(fmt.Sprintf("%s bt ON %s.%s = bt.%s", tableBannerTag, tableBanner, colId, colBannerID)).
		GroupBy(fmt.Sprintf("%s.%s", tableBanner, colId)).
		Where(sq.Eq{colId: bannerID}).
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
		&result.IsActive,
		&result.Content,
		&result.FeatureID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.TagIDs,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
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

func (b *BannerRepository) GetActiveByTagID(tagID int) ([]banner.Banner, error) {
	_ = tagID
	return nil, nil
}
