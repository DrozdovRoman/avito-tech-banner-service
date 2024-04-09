package v1

import (
	"encoding/json"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type BannerHandler struct {
	bannerService service.BannerService
}

func NewBannerHandler(bannerService *service.BannerService) *BannerHandler {
	return &BannerHandler{bannerService: *bannerService}
}

func (h *BannerHandler) GetBannerByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	banner, err := h.bannerService.GetByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banner)
}
