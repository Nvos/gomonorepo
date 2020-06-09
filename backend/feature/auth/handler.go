package auth

import (
	"backend/infra"
	"backend/security"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"time"
)

const refreshTokenCookie = "refresh"

type Dependencies struct {
	Cfg    *infra.Security
	Logger *infra.Logger
}

func RefreshTokenHandler(deps *Dependencies) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		session := c.Cookies(refreshTokenCookie)
		if session == "" {
			c.Next(security.ErrUnauthorized)
			return
		}

		claims := jwt.MapClaims{
			"id":   0,
			"name": "admin",
		}

		accessToken, err := security.GenerateAccessToken(deps.Cfg, claims)
		if err != nil {
			c.Next(fmt.Errorf("generate access token: %w", err))
			return
		}

		refreshToken, err := security.GenerateRefreshToken(deps.Cfg, claims)
		if err != nil {
			c.Next(fmt.Errorf("generate refresh token: %w", err))
			return
		}

		c.Cookie(&fiber.Cookie{
			HTTPOnly: true,
			Name:     refreshTokenCookie,
			Value:    refreshToken,
			// Secure = true, requires SSL in chrome, if set without it then cookie doesn't work. Works on FF thought
			Secure:   false,
			Expires:  time.Now().Add(time.Hour * 24 * 7),
			SameSite: "Lax",
		})

		err = c.JSON(fiber.Map{"accessToken": accessToken})
		if err != nil {
			c.Next(fmt.Errorf("json marshal: %w", err))
			return
		}
	}
}

func AccessTokenHandler(deps *Dependencies) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		var input AuthenticationInput
		err := c.BodyParser(&input)
		if err != nil {
			c.Next(err)
			return
		}

		if input.Password != "admin" || input.Username != "admin" {
			c.Next(security.ErrUnauthorized)
		}

		claims := jwt.MapClaims{
			"id":   0,
			"name": "admin",
		}

		accessToken, err := security.GenerateAccessToken(deps.Cfg, claims)
		if err != nil {
			c.Next(fmt.Errorf("generate access token: %w", err))
			return
		}

		refreshToken, err := security.GenerateRefreshToken(deps.Cfg, claims)
		if err != nil {
			c.Next(fmt.Errorf("generate refresh token: %w", err))
			return
		}

		c.Cookie(&fiber.Cookie{
			HTTPOnly: true,
			Name:     refreshTokenCookie,
			Value:    refreshToken,
			// Secure = true, requires SSL in chrome, if set without it then cookie doesn't work. Works on FF thought
			Secure:   false,
			Expires:  time.Now().Add(time.Hour * 24 * 7),
			SameSite: "Lax",
		})

		err = c.JSON(fiber.Map{"accessToken": accessToken})
		if err != nil {
			c.Next(fmt.Errorf("json marshal: %w", err))
			return
		}
	}
}

func MeHandler() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		token := c.Locals("user").(*jwt.Token)
		if token == nil {
			c.Next(security.ErrUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		id := claims["id"].(float64)

		err := c.JSON(&User{
			ID:       int(id),
			Username: name,
		})

		if err != nil {
			c.Next(err)
		}
	}
}
