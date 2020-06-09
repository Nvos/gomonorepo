package rest

import (
	"backend/infra"
	"backend/security"
	"errors"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"go.uber.org/zap"
)

func New(logger *infra.Logger) *fiber.App {
	router := applyMiddleware(fiber.New(), logger)
	router.Settings.ErrorHandler = ErrorHandler(logger)

	return router
}

func applyMiddleware(router *fiber.App, log *infra.Logger) *fiber.App {
	// Cors handling
	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000"},
	}))

	// Default logging
	router.Use(logger.New())

	// Default security headers
	router.Use(helmet.New())

	// Panic recovery middleware
	router.Use(recover.New(recover.Config{
		Handler: RecoveryHandler(log),
	}))

	return router
}

func ErrorHandler(logger *infra.Logger) func(ctx *fiber.Ctx, err error) {
	return func(ctx *fiber.Ctx, err error) {
		logger.Error("rest error handler", zap.Error(err))

		if errors.Is(err, security.ErrUnauthorized) {
			ctx.SendStatus(fiber.StatusUnauthorized)
			return
		}

		ctx.SendStatus(fiber.StatusInternalServerError)
	}
}

func RecoveryHandler(logger *infra.Logger) func(c *fiber.Ctx, err error) {
	return func(c *fiber.Ctx, err error) {
		logger.Error("rest recovery handler", zap.Error(err))
		c.SendStatus(500)
	}
}
