package v1

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/dependency"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

// AuthenticationAndAuthorizationHandler handle user authentication and authorization here.
func AuthenticationAndAuthorizationHandler(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) (err error) {
		if config.EnableAuthentication {
			bearerToken:=context.Request().Header.Get("Authorization")
			if bearerToken==""{
				return common.GenerateErrorResponse(context,"[ERROR]: No token found!","Please provide a valid token!")
			}
			var token string
			if len(strings.Split(bearerToken," "))==2{
				token=strings.Split(bearerToken," ")[1]
			}else{
				return common.GenerateErrorResponse(context,"[ERROR]: No token found!","Please provide a valid token!")
			}
			res, _ := dependency.GetV1JwtService().ValidateToken(token)
			if !res {
				return common.GenerateErrorResponse(context, "[ERROR]: Invalid token!", "Please provide a valid token!")
			}
		}
		return handler(context)
	}
}

// AuthenticationHandler handles user authentication.
func AuthenticationHandler(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) (err error) {
		if config.EnableAuthentication {
			bearerToken:=context.Request().Header.Get("Authorization")
			if bearerToken==""{
				return common.GenerateUnauthorizedResponse(context,"[ERROR]: No token found!","Please provide a valid token!")
			}
			var token string
			if len(strings.Split(bearerToken," "))==2{
				token=strings.Split(bearerToken," ")[1]
			}else{
				return common.GenerateUnauthorizedResponse(context,"[ERROR]: No token found!","Please provide a valid token!")
			}
			res, _ := dependency.GetV1JwtService().ValidateToken(token)
			if !res {
				return common.GenerateUnauthorizedResponse(context, "[ERROR]: Invalid token!", "Please provide a valid token!")
			}
		}
		return handler(context)
	}
}

// AuthorizationHandler handles user authorization.
func AuthorizationHandler(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) (err error) {
		if config.EnableAuthentication {
			bearerToken:=context.Request().Header.Get("Authorization")
			if bearerToken==""{
				return common.GenerateErrorResponse(context,"[ERROR]: No token found!","Please provide a valid token!")
			}
			var token string
			if len(strings.Split(bearerToken," "))==2{
				token=strings.Split(bearerToken," ")[1]
			}else{
				return common.GenerateErrorResponse(context,"[ERROR]: No token found!","Please provide a valid token!")
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

		}
		return handler(context)
	}
}