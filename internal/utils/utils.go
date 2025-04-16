package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ElderLab/CrazyLabelling"
	"github.com/ElderLab/SSO-Middleware/Claims"
	"github.com/ElderLab/SSO-Middleware/User"
	"github.com/gofiber/fiber/v2"
)

func QuerySSO(token string) (bool, Claims.SSOClaims) {
	finalURL := fmt.Sprintf(CrazyLabelling.ApiInternalProtocol+CrazyLabelling.SSOBack+"/validate?token=%s", token)
	// send a get request using fiber
	agent := fiber.Get(finalURL)
	status := fiber.StatusInternalServerError
	var body []byte
	var errs []error
	status, body, errs = agent.Bytes()
	if len(errs) != 0 {
		return false, Claims.SSOClaims{}
	}
	if status == fiber.StatusOK {
		var ssoClaims Claims.SSOClaims
		err := json.Unmarshal(body, &ssoClaims)
		if err != nil {
			return false, Claims.SSOClaims{}
		}
		return true, ssoClaims
	}
	return false, Claims.SSOClaims{}
}

func GetFullUser(claims Claims.SSOClaims) (error, User.User) {
	finalURL := fmt.Sprintf(CrazyLabelling.ApiInternalProtocol+CrazyLabelling.AccessDomain+"/api/sso/fulluser?username=%s", claims.Username)
	// send a get request using fiber
	agent := fiber.Get(finalURL)
	status := fiber.StatusInternalServerError
	var body []byte
	var errs []error
	status, body, errs = agent.Bytes()
	if len(errs) != 0 {
		return errs[0], User.User{}
	}
	if status == fiber.StatusOK {
		var fullUser User.User
		err := json.Unmarshal(body, &fullUser)
		if err != nil {
			return err, User.User{}
		}
		return nil, fullUser
	}
	return nil, User.User{}
}
