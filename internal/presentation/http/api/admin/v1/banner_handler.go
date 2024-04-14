package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/service"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api/common/dto"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AdminBannerHandler struct {
	bannerService service.BannerService
}

func NewAdminBannerHandler(bannerService *service.BannerService) *AdminBannerHandler {
	return &AdminBannerHandler{bannerService: *bannerService}
}

func (h *AdminBannerHandler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateBannerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	logrus.Info(req)

	ctx := r.Context()
	bannerID, err := h.bannerService.CreateBanner(ctx, req.TagIDs, req.FeatureID, req.Content, req.IsActive)
	if err != nil {
		switch {
		case errors.Is(err, banner.ErrNoFeatureID), errors.Is(err, banner.ErrNoTagIDs), errors.Is(err, banner.ErrJSONMarshal):
			http.Error(w, fmt.Sprintf("validation error: %v", err), http.StatusBadRequest)
		default:
			http.Error(w, fmt.Sprintf("internal server error: %v", err), http.StatusInternalServerError)
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"banner_id": bannerID})
}
