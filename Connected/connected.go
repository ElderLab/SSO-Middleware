package Connected

import (
	"github.com/ElderLab/SSO-Middleware/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Filter    func(c *fiber.Ctx) bool
	GetClaims bool
	GetUser   bool
}

var ConfigDefault = Config{
	Filter:    nil,
	GetClaims: false,
	GetUser:   false,
}

func configDefault(config ...Config) Config {
	return ConfigDefault
}

// Middleware that checks if the user is connected to the SSO service, you can find the claims and the user information in Locals named "claims" and "user" respectively
func New(config Config) fiber.Handler {
	// Return new middleware
	return func(c *fiber.Ctx) error {
		// If Filter is not nil and returns true, skip middleware
		if config.Filter != nil && config.Filter(c) {
			return c.Next()
		}
		// get the token from bearer token
		token := c.Get("Authorization")
		// if the token is empty, return unauthorized
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		token = token[7:]
		// query the SSO service
		connected, claims := utils.QuerySSO(token)
		// if the SSO service is not connected, return unauthorized
		if !connected {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		// if GetClaims is true, set the claims to the context
		if config.GetClaims {
			c.Locals("claims", claims)
		}
		if config.GetUser {
			// get the full user information
			err, fullUser := utils.GetFullUser(claims)
			// if there is an error, return internal server error
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
			}
			// set the full user information to the context
			c.Locals("user", fullUser)
		}
		// Continue stack
		return c.Next()
	}
}
