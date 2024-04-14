package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/service"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/presentation/http/api/common/dto/banner_dto"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type AdminBannerHandler struct {
	bannerService service.BannerService
}

func NewAdminBannerHandler(bannerService *service.BannerService) *AdminBannerHandler {
	return &AdminBannerHandler{bannerService: *bannerService}
}

func (h *AdminBannerHandler) GetBanners(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var featureID, tagID, limit, offset int
	var err error

	if fid := r.URL.Query().Get("feature_id"); fid != "" {
		var err error
		featureID, err = strconv.Atoi(fid)
		if err != nil {
			http.Error(w, "Invalid feature ID", http.StatusBadRequest)
			return
		}
	}

	if tid := r.URL.Query().Get("tag_id"); tid != "" {
		var err error
		tagID, err = strconv.Atoi(tid)
		if err != nil {
			http.Error(w, "Invalid tag ID", http.StatusBadRequest)
			return
		}
	}

	if l := r.URL.Query().Get("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		} else {
			http.Error(w, "Invalid limit value", http.StatusBadRequest)
			return
		}
	} else {
		limit = 10
	}

	if o := r.URL.Query().Get("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil && val >= 0 {
			offset = val
		} else {
			http.Error(w, "Invalid offset value", http.StatusBadRequest)
			return
		}
	}

	banners, err := h.bannerService.GetBanners(ctx, tagID, featureID, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses, err := banner_dto.NewBannerResponsesFromDomain(banners)
	if err != nil {
		http.Error(w, "Failed to process banners", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses)
}

func (h *AdminBannerHandler) CreateBanner(w http.ResponseWriter, r *http.Request) {
	var req banner_dto.CreateBannerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	bannerID, err := h.bannerService.CreateBanner(r.Context(), req.TagIDs, req.FeatureID, req.Content, req.IsActive)
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

func (h *AdminBannerHandler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bannerID, err := strconv.Atoi(chi.URLParam(r, "banner_id"))
	if err != nil {
		http.Error(w, "invalid banner ID", http.StatusBadRequest)
		return
	}

	err = h.bannerService.DeleteBanner(ctx, bannerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *AdminBannerHandler) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bannerID, err := strconv.Atoi(chi.URLParam(r, "banner_id"))
	if err != nil {
		http.Error(w, "invalid banner ID", http.StatusBadRequest)
		return
	}

	var req banner_dto.PatchBannerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	err = h.bannerService.UpdateBanner(ctx, bannerID, req.TagIDs, req.FeatureID, req.Content, req.IsActive)
	if err != nil {
		if errors.Is(err, banner.ErrNoFeatureID) || errors.Is(err, banner.ErrNoTagIDs) || errors.Is(err, banner.ErrJSONMarshal) {
			http.Error(w, fmt.Sprintf("validation error: %v", err), http.StatusBadRequest)
			return
		}
		if errors.Is(err, banner.ErrBannerNotFound) {
			http.Error(w, fmt.Sprintf("banner not found: %v", err), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
