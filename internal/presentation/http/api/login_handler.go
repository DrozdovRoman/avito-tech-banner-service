package api

import (
	"encoding/json"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LoginHandler struct {
	UserService service.UserService
}

var creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewLoginHandler(userService *service.UserService) *LoginHandler {
	return &LoginHandler{UserService: *userService}
}

func (l *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		logrus.WithError(err).Error("Failed to decode JSON body")
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}

	logrus.WithField("username", creds.Username).Info("Login attempt")

	token, err := l.UserService.Authenticate(creds.Username, creds.Password)
	if err != nil {
		logrus.WithError(err).Error("Failed to authenticate")
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
