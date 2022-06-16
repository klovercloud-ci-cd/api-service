package v1

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

func checkAuthority(userResourcePermission v1.UserResourcePermission, resourceName, role, permission string) error {
	var resourceWiseRoles v1.ResourceWiseRoles
	for _, resource := range userResourcePermission.Resources {
		if resource.Name == resourceName {
			resourceWiseRoles = resource
			break
		}
	}
	if role != "" {
		for _, each := range resourceWiseRoles.Roles {
			if each.Name == role {
				return nil
			}
		}
	} else if permission != "" {
		for _, each := range resourceWiseRoles.Roles {
			for _, perm := range each.Permissions {
				if perm.Name == permission {
					return nil
				}
			}

		}
	}
	return errors.New("[ERROR]: Insufficient permission")
}

func GetUserResourcePermissionFromBearerToken(context echo.Context, jwtService service.Jwt) (v1.UserResourcePermission, error) {
	bearerToken := context.Request().Header.Get("Authorization")
	if bearerToken == "" {
		bearerToken = context.QueryParam("token")
		if bearerToken == "" {
			return v1.UserResourcePermission{}, errors.New("[ERROR]: No token found")
		}
		bearerToken="Bearer "+bearerToken
	}
	var token string
	if len(strings.Split(bearerToken, " ")) == 2 {
		token = strings.Split(bearerToken, " ")[1]
	} else {
		return v1.UserResourcePermission{}, errors.New("[ERROR]: No token found")
	}
	res, _ := jwtService.ValidateToken(token)
	if !res {
		return v1.UserResourcePermission{}, errors.New("[ERROR]: Token is expired")
	}
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	jsonBody, err := json.Marshal(claims["data"])
	if err != nil {
		log.Println(err)
	}
	userResourcePermission := v1.UserResourcePermission{}
	if err := json.Unmarshal(jsonBody, &userResourcePermission); err != nil {
		return v1.UserResourcePermission{}, errors.New("[ERROR]: No resource permissions")
	}
	return userResourcePermission, nil
}

func GetClientNameFromToken(context echo.Context, jwtService service.Jwt) (v1.AgentData, error) {
	token := context.Request().Header.Get("token")
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	jsonBody, err := json.Marshal(claims["data"])
	if err != nil {
		log.Println(err.Error())
	}
	agent := v1.AgentData{}
	if err := json.Unmarshal(jsonBody, &agent); err != nil {
		return v1.AgentData{}, errors.New("[ERROR]: No resource permissions")
	}
	return agent, nil
}
