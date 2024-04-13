package jwt

import (
	"fmt"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/configuration"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type JWTUtils struct {
	key []byte
}

func NewJWTUtils(config *configuration.Configuration) *JWTUtils {
	return &JWTUtils{key: []byte(config.SecretKey)}
}

func (j *JWTUtils) GenerateToken(username string, isAdmin bool) (string, error) {
	claims := &Claims{
		Username: username,
		Admin:    isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
			Issuer:    "bannerService",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.key)
}

func (j *JWTUtils) ValidateToken(authHeader string) (*Claims, error) {
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return nil, fmt.Errorf("invalid Authorization token")
	}
	tokenStr := splitToken[1]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorExpired)
	}
	return claims, nil
}
