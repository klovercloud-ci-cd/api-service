package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/dependency"
	"github.com/labstack/echo/v4"
	"strings"
)


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
