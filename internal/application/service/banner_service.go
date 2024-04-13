package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/cache"
	"time"
)

type BannerService struct {
	bannerRepo banner.Repository
	cache      cache.Cache
}

func NewBannerService(bannerRepo banner.Repository, cache cache.Cache) *BannerService {
	return &BannerService{bannerRepo: bannerRepo, cache: cache}
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

func (b *BannerService) GetByID(ctx context.Context, id int) (banner.Banner, error) {
	return b.bannerRepo.GetByID(ctx, id)
}

func (b *BannerService) GetAll() ([]banner.Banner, error) {
	return b.bannerRepo.GetAll()
}
