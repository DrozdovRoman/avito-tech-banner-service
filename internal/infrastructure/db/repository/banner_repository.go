package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/db"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

const (
	tableTag = "tag"
	colTagID = "id"

	tableFeature = "feature"
	colFeatureID = "id"

	tableBanner    = "banner"
	colId          = "id"
	colIsActive    = "is_active"
	colContent     = "content"
	colFKFeatureID = "feature_id"
	colCreatedAt   = "created_at"
	colUpdatedAt   = "updated_at"

	tableBannerFeatureTag = "banner_feature_tag"
	colFKTagID            = "tag_id"
	colFKBannerID         = "banner_id"
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

func (b *BannerRepository) GetActiveBannerContentByTagAndFeature(ctx context.Context, tagID int, featureID int) (json.RawMessage, error) {

	builderSelectActiveBannerContentByTagAndFeature := sq.Select(tableBanner + "." + colContent).
		From(tableBanner).
		Join(fmt.Sprintf("%s ON %s.%s = %s.%s",
			tableBannerFeatureTag, tableBanner, colFKFeatureID, tableBannerFeatureTag, colFKFeatureID),
		).
		Where(sq.And{
			sq.Eq{tableBanner + "." + colIsActive: true},
			sq.Eq{tableBanner + "." + colFKFeatureID: featureID},
			sq.Eq{tableBannerFeatureTag + "." + colFKTagID: tagID},
		}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderSelectActiveBannerContentByTagAndFeature.ToSql()

	if err != nil {
		return json.RawMessage{}, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "banner_repository.GetActiveBannerContentByTagAndFeature",
		QueryRaw: query,
	}

	var result json.RawMessage

	err = b.db.DB().QueryRowContext(ctx, q, args...).Scan(&result)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return json.RawMessage{}, banner.ErrBannerNotFound
		}
		return json.RawMessage{}, fmt.Errorf("failed to scan row: %v", err)
	}

	return result, nil
}

func (b *BannerRepository) AddBanner(ctx context.Context, banner *banner.Banner) (int, error) {
	builderInsertBanner := sq.Insert(tableBanner).
		Columns(colIsActive, colContent, colFKFeatureID, colUpdatedAt, colCreatedAt).
		Values(banner.IsActive, banner.Content, banner.FeatureID, banner.UpdatedAt, banner.CreatedAt).
		Suffix("RETURNING " + colId).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := builderInsertBanner.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build insert SQL query: %v", err)
	}

	q := db.Query{
		Name:     "banner_repository.Insert",
		QueryRaw: sql,
	}

	var newID int
	err = b.db.DB().QueryRowContext(ctx, q, args...).Scan(&newID)
	if err != nil {
		return 0, fmt.Errorf("failed to scan row: %v", err)
	}

	if len(banner.TagIDs) > 0 {
		if banner.FeatureID.Valid { // Only proceed if the FeatureID is not NULL
			err = b.addBannerFeatureTags(ctx, newID, int(banner.FeatureID.Int64), banner.TagIDs)
			if err != nil {
				return 0, fmt.Errorf("failed to add banner tags: %v", err)
			}
		} else {
			return 0, fmt.Errorf("feature ID is null")
		}
	}

	return newID, nil
}

func (b *BannerRepository) addBannerFeatureTags(ctx context.Context, bannerID int, featureID int, tagIDs []int) error {
	for _, tagID := range tagIDs {
		logrus.Info(tagID)
		sql, args, err := sq.Insert(tableBannerFeatureTag).
			Columns(colFKBannerID, colFKFeatureID, colFKTagID).
			Values(bannerID, featureID, tagID).
			PlaceholderFormat(sq.Dollar).
			ToSql()

		if err != nil {
			return fmt.Errorf("failed to build insert tag SQL query: %v", err)
		}

		query := db.Query{
			Name:     "insert_tag_association",
			QueryRaw: sql,
		}

		_, err = b.db.DB().ExecContext(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("failed to execute insert tag query: %v", err)
		}
	}

	return nil
}
