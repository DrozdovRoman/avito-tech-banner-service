package service

import "github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/jwt"

type UserService struct {
	JWTUtils *jwt.JWTUtils
}

func NewUserService(jwtUtils *jwt.JWTUtils) *UserService {
	return &UserService{JWTUtils: jwtUtils}
}

func (s *UserService) Authenticate(username, password string) (string, error) {
	isAdmin := password == "avitoDeveloper"
	return s.JWTUtils.GenerateToken(username, isAdmin)
}

func (s *UserService) ValidateToken(token string) (*jwt.Claims, error) {
	return s.JWTUtils.ValidateToken(token)
}
