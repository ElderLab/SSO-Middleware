package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ElderLab/CrazyLabelling"
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

func QuerySSO(token string) (bool, SSOClaims) {
	finalURL := fmt.Sprintf(CrazyLabelling.ApiInternalProtocol+CrazyLabelling.SSOBack+"/validate?token=%s", token)
	// send a get request using fiber
	agent := fiber.Get(finalURL)
	status := fiber.StatusInternalServerError
	var body []byte
	var errs []error
	status, body, errs = agent.Bytes()
	if len(errs) != 0 {
		return false, SSOClaims{}
	}
	if status == fiber.StatusOK {
		var ssoClaims SSOClaims
		err := json.Unmarshal(body, &ssoClaims)
		if err != nil {
			return false, SSOClaims{}
		}
		return true, ssoClaims
	}
	return false, SSOClaims{}
}

type FullUser struct {
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Roles    []Role `json:"roles"`
}

func GetFullUser(claims SSOClaims) (error, FullUser) {
	finalURL := fmt.Sprintf(CrazyLabelling.ApiInternalProtocol+CrazyLabelling.AccessDomain+"/fulluser?username=%s", claims.Username)
	// send a get request using fiber
	agent := fiber.Get(finalURL)
	status := fiber.StatusInternalServerError
	var body []byte
	var errs []error
	status, body, errs = agent.Bytes()
	if len(errs) != 0 {
		return errs[0], FullUser{}
	}
	if status == fiber.StatusOK {
		var fullUser FullUser
		err := json.Unmarshal(body, &fullUser)
		if err != nil {
			return err, FullUser{}
		}
		return nil, fullUser
	}
	return nil, FullUser{}
}
