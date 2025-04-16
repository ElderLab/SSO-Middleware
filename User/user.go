package User

import (
	"github.com/ElderLab/SSO-Middleware/Claims"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       int           `json:"id"`
	UUID     string        `json:"uuid"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Roles    []Claims.Role `json:"roles"`
}

func GetUser(c *fiber.Ctx) (User, error) {
	val := c.Locals("user")
	if val == nil {
		return User{}, fiber.NewError(fiber.StatusInternalServerError, "no full user found")
	}
	fullUser, ok := val.(User)
	if !ok {
		return User{}, fiber.NewError(fiber.StatusInternalServerError, "invalid full user type")
	}
	return fullUser, nil
}
