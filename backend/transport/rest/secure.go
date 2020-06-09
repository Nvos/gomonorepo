package rest

import (
	"backend/infra"
	"github.com/gofiber/fiber"
	"github.com/gofiber/jwt" // jwtware
)

func SecurityMiddleware(cfg *infra.Security, logger *infra.Logger) func(*fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(cfg.JwtSigningSecret),
		SigningMethod: "HS256",
		TokenLookup:   "header:authorization",
	})
}
