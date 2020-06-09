package security

import (
	"backend/infra"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var ErrUnauthorized = errors.New("unauthorized")

func GenerateAccessToken(cfg *infra.Security, claims jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	claims["exp"] = time.Now().Add(time.Duration(cfg.JwtTokenLifetime) * time.Millisecond).Unix()

	return token.SignedString([]byte(cfg.JwtSigningSecret))
}

func GenerateRefreshToken(cfg *infra.Security, claims jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	claims["exp"] = time.Now().Add(time.Duration(cfg.JwtRefreshTokenLifetime) * time.Millisecond).Unix()

	return token.SignedString([]byte(cfg.JwtSigningSecret))
}
