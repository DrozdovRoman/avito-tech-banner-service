package service

import (
	"context"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
)

type BannerService struct {
	bannerRepo banner.Repository
}

func NewBannerService(bannerRepo banner.Repository) *BannerService {
	return &BannerService{bannerRepo: bannerRepo}
}

func (b *BannerService) GetByID(ctx context.Context, id int) (banner.Banner, error) {
	return b.bannerRepo.GetByID(ctx, id)
}
