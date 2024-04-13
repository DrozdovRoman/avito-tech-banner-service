package v1

import (
	"encoding/json"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/service"
	"net/http"
)

type AdminBannerHandler struct {
	bannerService service.BannerService
}

func NewAdminBannerHandler(bannerService *service.BannerService) *AdminBannerHandler {
	return &AdminBannerHandler{bannerService: *bannerService}
}

func (h *AdminBannerHandler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	var req CreateBannerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	bannerId, err := h.bannerService.CreateBanner(ctx, req.TagIDs, req.FeatureID, req.Content, req.IsActive)
}
