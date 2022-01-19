package v1

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

func getResourceWiseRoleFromToken(context echo.Context, resource string)*v1.ResourceWiseRoles{
	bearerToken:=context.Request().Header.Get("Authorization")
	if bearerToken==""{
		return nil
	}
	var token string
	if len(strings.Split(bearerToken," "))==2{
		token=strings.Split(bearerToken," ")[1]
	}else{
		return nil
	}
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	jsonbody, err := json.Marshal(claims["data"])
	if err != nil {
		log.Println(err)
	}
	usersPermission := v1.UserResourcePermission{}
	if err := json.Unmarshal(jsonbody, &usersPermission); err != nil {
		log.Println(err)
	}
	for _,each:=range usersPermission.Resources{
		if each.Name==resource{
			return &each
		}
	}
	return nil
}