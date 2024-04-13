package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/service"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"net/http"
	"strconv"
)

type UserBannerHandler struct {
	bannerService service.BannerService
}

func NewUserBannerHandler(bannerService *service.BannerService) *UserBannerHandler {
	return &UserBannerHandler{bannerService: *bannerService}
}

func (h *UserBannerHandler) FetchActiveUserBannerContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var featureID, tagID int
	var err error

	tagID, err = strconv.Atoi(r.URL.Query().Get("tag_id"))
	if err != nil {
		h.respondWithError(w, "tag_id is required and must be a numeric value.", http.StatusBadRequest)
		return
	}

	featureID, err = strconv.Atoi(r.URL.Query().Get("feature_id"))
	if err != nil {
		h.respondWithError(w, "feature_id is required and must be a numeric value.", http.StatusBadRequest)
		return
	}

	useLastRevision := r.URL.Query().Get("use_last_revision") == "true"

	content, err := h.bannerService.GetUserBannerActiveContent(ctx, tagID, featureID, useLastRevision)
	if err != nil {
		if errors.Is(err, banner.ErrBannerNotFound) {
			h.respondWithError(w, "The specified banner does not exist.", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`"%s"`, content)))
}

func (h *UserBannerHandler) respondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := map[string]string{"error": message}
	json.NewEncoder(w).Encode(errorResponse)
}
