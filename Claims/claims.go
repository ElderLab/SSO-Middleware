package Claims

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type SSOClaims struct {
	UUID     string `json:"UUID"`
	Username string `json:"Username"`
	Roles    []Role `json:"Roles"`
}

type Role struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func GetClaims(c *fiber.Ctx) (SSOClaims, error) {
	val := c.Locals("claims")
	if val == nil {
		return SSOClaims{}, errors.New("no claims found")
	}
	claims, ok := val.(SSOClaims)
	if !ok {
		return SSOClaims{}, errors.New("invalid claims type")
	}
	return claims, nil
}
