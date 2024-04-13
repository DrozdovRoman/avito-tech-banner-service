package v1

import (
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/service"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	featureID, err = strconv.Atoi(r.URL.Query().Get("feature_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	useLastRevision := r.URL.Query().Get("use_last_revision") == "true"

	content, err := h.bannerService.GetUserBannerActiveContent(ctx, tagID, featureID, useLastRevision)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}
