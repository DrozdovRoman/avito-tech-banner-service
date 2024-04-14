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

func (b *BannerRepository) GetBanner(ctx context.Context, id int) (*banner.Banner, error) {
	builderSelectBanner := sq.Select(
		"b."+colId, "b."+colIsActive, "b."+colContent, "b."+colFKFeatureID, "b."+colCreatedAt, "b."+colUpdatedAt,
		fmt.Sprintf("COALESCE(array_agg(bft.tag_id) FILTER (WHERE bft.tag_id IS NOT NULL), '{}') AS tag_ids")).
		From(fmt.Sprintf("%s AS b", tableBanner)).
		LeftJoin(fmt.Sprintf("%s AS bft ON b.%s = bft.%s", tableBannerFeatureTag, colId, colFKBannerID)).
		GroupBy(fmt.Sprintf("b.%s", colId)).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := builderSelectBanner.Where(sq.Eq{colId: id}).ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "banner_repository.GetBanner",
		QueryRaw: sql,
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
			return nil, banner.ErrBannerNotFound
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &result, nil
}

func (b *BannerRepository) GetBanners(ctx context.Context, tagID, featureID, limit, offset int) ([]banner.Banner, error) {
	builderSelectBanners := sq.Select(
		"b."+colId, "b."+colIsActive, "b."+colContent, "b."+colFKFeatureID, "b."+colCreatedAt, "b."+colUpdatedAt,
		fmt.Sprintf("COALESCE(array_agg(bft.tag_id) FILTER (WHERE bft.tag_id IS NOT NULL), '{}') AS tag_ids")).
		From(fmt.Sprintf("%s AS b", tableBanner)).
		LeftJoin(fmt.Sprintf("%s AS bft ON b.%s = bft.%s", tableBannerFeatureTag, colId, colFKBannerID)).
		GroupBy(fmt.Sprintf("b.%s", colId)).
		PlaceholderFormat(sq.Dollar)

	if featureID != 0 {
		builderSelectBanners = builderSelectBanners.Where(sq.Eq{"b.feature_id": featureID})
	}

	if tagID != 0 {
		builderSelectBanners = builderSelectBanners.Where(sq.Eq{"bft.tag_id": tagID})
	}

	builderSelectBanners = builderSelectBanners.
		Limit(uint64(limit)).
		Offset(uint64(offset))

	query, args, err := builderSelectBanners.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build insert SQL query: %v", err)
	}

	q := db.Query{
		Name:     "banner_repository.GetBanners",
		QueryRaw: query,
	}

	var banners []banner.Banner
	err = b.db.DB().ScanAllContext(ctx, &banners, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return banners, nil
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

	if len(banner.TagIDs) > 0 { // Only proceed if the FeatureID is not NULL
		err = b.addBannerFeatureTags(ctx, newID, banner.FeatureID, banner.TagIDs)
		if err != nil {
			return 0, fmt.Errorf("failed to add banner tags: %v", err)
		}
	}

	return newID, nil
}

func (b *BannerRepository) UpdateBanner(ctx context.Context, banner *banner.Banner) error {
	builderUpdate := sq.Update(tableBanner).
		Set(colIsActive, banner.IsActive).
		Set(colContent, banner.Content).
		Set(colFKFeatureID, banner.FeatureID).
		Set(colUpdatedAt, banner.UpdatedAt).
		Where(sq.Eq{colId: banner.ID}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := builderUpdate.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build insert SQL query: %v", err)
	}

	q := db.Query{
		Name:     "banner_repository.Update",
		QueryRaw: sql,
	}

	res, err := b.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to update banner: %v tag: %v", err, res)
	}

	err = b.updateBannerFeatureTag(ctx, banner)
	if err != nil {
		return fmt.Errorf("failed to update banner feature tag: %v", err)
	}

	return nil
}

func (b *BannerRepository) updateBannerFeatureTag(ctx context.Context, banner *banner.Banner) error {
	builderDelete := sq.Delete(tableBannerFeatureTag).
		Where(sq.Eq{colFKBannerID: banner.ID}).
		PlaceholderFormat(sq.Dollar)

	sqlDelete, argsDelete, err := builderDelete.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build delete SQL query: %v", err)
	}

	qDelete := db.Query{
		Name:     "banner_repository.DeleteBannerFeatureTag",
		QueryRaw: sqlDelete,
	}

	_, err = b.db.DB().ExecContext(ctx, qDelete, argsDelete...)
	if err != nil {
		return fmt.Errorf("failed to delete banner feature tag: %v", err)
	}

	for _, tagID := range banner.TagIDs {
		builderInsert := sq.Insert(tableBannerFeatureTag).
			Columns(colFKFeatureID, colFKTagID, colFKBannerID).
			Values(banner.FeatureID, tagID, banner.ID).
			PlaceholderFormat(sq.Dollar)

		sqlInsert, argsInsert, err := builderInsert.ToSql()
		if err != nil {
			return fmt.Errorf("failed to build insert SQL query: %v", err)
		}

		qInsert := db.Query{
			Name:     "banner_repository.InsertBannerFeatureTag",
			QueryRaw: sqlInsert,
		}

		_, err = b.db.DB().ExecContext(ctx, qInsert, argsInsert...)
		if err != nil {
			return fmt.Errorf("failed to insert banner feature tag: %v", err)
		}
	}

	return nil
}

func (b *BannerRepository) addBannerFeatureTags(ctx context.Context, bannerID int, featureID int, tagIDs []int) error {
	for _, tagID := range tagIDs {
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

func (b *BannerRepository) DeleteBanner(ctx context.Context, id int) error {
	builderDeleteBannner := sq.Delete(tableBanner).
		Where(sq.Eq{colId: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderDeleteBannner.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build insert SQL query: %v", err)
	}

	q := db.Query{
		Name:     "banner_repository.Delete",
		QueryRaw: query,
	}

	_, err = b.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("failed to execute insert query: %v", err)
	}

	return nil
}

func (b *BannerRepository) BannerExists(ctx context.Context, id int) (bool, error) {
	var exist bool
	query, args, err := sq.Select(colId).
		From(tableBanner).
		Where(sq.Eq{colId: id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "banner_repository.Exists",
		QueryRaw: query,
	}

	err = b.db.DB().QueryRowContext(ctx, q, args...).Scan(&exist)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("failed to scan row: %v", err)
	}

	return exist, nil
}
