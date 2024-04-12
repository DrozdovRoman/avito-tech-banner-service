package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

// Claims extends the standard jwt.RegisteredClaims with custom claim information.
type Claims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}
