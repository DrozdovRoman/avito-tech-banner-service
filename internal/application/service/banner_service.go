package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/cache"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/db"
	"time"
)

type BannerService struct {
	bannerRepo banner.Repository
	cache      cache.Cache
	txManager  db.TxManager
}

func NewBannerService(
	bannerRepo banner.Repository, cache cache.Cache, txManager db.TxManager) *BannerService {
	return &BannerService{bannerRepo: bannerRepo, cache: cache, txManager: txManager}
}

func (b *BannerService) GetUserBannerActiveContent(ctx context.Context, tagID, featureID int, useLastVersion bool) (json.RawMessage, error) {
	cacheKey := fmt.Sprintf("banner_%d_%d", featureID, tagID)

	if cachedContent, found := b.cache.Get(cacheKey); !useLastVersion && found {
		return cachedContent.(json.RawMessage), nil
	}

	bannerContent, err := b.bannerRepo.GetActiveBannerContentByTagAndFeature(ctx, tagID, featureID)
	if err != nil {
		return nil, err
	}

	b.cache.Set(cacheKey, bannerContent, 5*time.Minute)

	return bannerContent, nil
}

func (b *BannerService) GetBanners(ctx context.Context, tagID, featureID, limit, offset int) ([]banner.Banner, error) {
	return b.bannerRepo.GetBanners(ctx, tagID, featureID, limit, offset)
}

func (b *BannerService) GetBanner(ctx context.Context, id int) (*banner.Banner, error) {
	return b.bannerRepo.GetBanner(ctx, id)
}

func (b *BannerService) CreateBanner(ctx context.Context, tagIDs []int, featureID int, content string, isActive bool) (int, error) {
	newBanner, err := banner.NewBanner(tagIDs, featureID, content, isActive)
	if err != nil {
		return 0, err
	}

	var newID int
	err = b.txManager.ReadCommitted(ctx, func(txCtx context.Context) error {
		var errTx error
		newID, errTx = b.bannerRepo.AddBanner(txCtx, newBanner)
		if errTx != nil {
			return fmt.Errorf("failed to add banner: %w", errTx)
		}

		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("transaction failed: %w", err)
	}

	return newID, nil
}

func (b *BannerService) DeleteBanner(ctx context.Context, id int) error {
	existBanner, err := b.bannerRepo.BannerExists(ctx, id)
	if err != nil {
		return err
	}

	if !existBanner {
		return banner.ErrBannerNotFound
	}

	err = b.txManager.ReadCommitted(ctx, func(txCtx context.Context) error {
		return b.bannerRepo.DeleteBanner(txCtx, id)
	})

	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}

func (b *BannerService) UpdateBanner(ctx context.Context, id int, tagIDs []int, featureID int, content string, isActive bool) error {
	existingBanner, err := b.bannerRepo.GetBanner(ctx, id)
	if err != nil {
		return err
	}

	var changes bool

	if tagIDs != nil && len(tagIDs) > 0 {
		existingBanner.SetTagIDs(tagIDs)
		changes = true
	}

	if featureID != 0 && featureID != existingBanner.GetFeature() {
		existingBanner.SetFeature(featureID)
		changes = true
	}

	if isActive != existingBanner.GetIsActive() {
		existingBanner.SetIsActive(isActive)
		changes = true
	}

	if content != "" {
		existingContent := existingBanner.GetContent()
		contentJson, err := json.Marshal(content)
		if err != nil {
			return err
		}

		if !bytes.Equal(contentJson, existingContent) {
			fmt.Println("New content is different, updating...")
			existingBanner.SetContent(contentJson)
			changes = true
		}
	}

	if !changes {
		return nil
	}

	err = b.txManager.ReadCommitted(ctx, func(txCtx context.Context) error {
		return b.bannerRepo.UpdateBanner(txCtx, existingBanner)
	})

	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}
