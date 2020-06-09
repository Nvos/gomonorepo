package rest

import (
	"backend/feature/auth"
	"backend/infra"
	"github.com/gofiber/fiber"
)

func AuthGroup(router *fiber.App, logger *infra.Logger, cfg *infra.Security) *fiber.Group {
	group := router.Group("/auth")

	group.Post("/token", auth.AccessTokenHandler(&auth.Dependencies{
		Logger: logger,
		Cfg:    cfg,
	}))

	group.Post("/refresh", auth.RefreshTokenHandler(&auth.Dependencies{
		Logger: logger,
		Cfg:    cfg,
	}))

	group.Use(SecurityMiddleware(cfg, logger))
	group.Get("/me", auth.MeHandler())
	return group
}
