package middlewares

import (
	"context"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/jwt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/service"
	"net/http"
)

type AuthMiddleware struct {
	userService service.UserService
}

func NewAuthMiddleware(userService *service.UserService) *AuthMiddleware {
	return &AuthMiddleware{userService: *userService}
}

func (a *AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		claims, err := a.userService.ValidateToken(authHeader)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Добавляем claims в context
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *AuthMiddleware) AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value("claims").(*jwt.Claims)
		if !ok || !claims.Admin {
			http.Error(w, "Access denied: Admin only <3", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
