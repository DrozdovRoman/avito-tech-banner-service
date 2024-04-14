package service

import (
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

//func (b *BannerService) DeleteBanner(ctx context.Context, id int) error {
//	banner, err := b.bannerRepo.GetByID(ctx, id)
//}

func (b *BannerService) GetByID(ctx context.Context, id int) (banner.Banner, error) {
	return banner.Banner{}, nil
}

func (b *BannerService) GetAll() ([]banner.Banner, error) {
	return nil, nil
}
