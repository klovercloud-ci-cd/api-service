package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
)

type agentApi struct {
	agentService service.Agent
	jwtService       service.Jwt
}

func (a agentApi) Save(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	agentName:=context.QueryParam("name")
	if config.EnableAuthentication {
		agentObjFromToken, err := GetClientNameFromToken(context,a.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		agentName=agentObjFromToken.Name
	}
	code, err := a.agentService.Store(formData,agentName)
	if err != nil {
		return common.GenerateErrorResponse(context, err, err.Error())
	}
	return context.JSON(code,err)
}

func (a agentApi) Get(context echo.Context) error {
	agentName:=context.Param("name")
	if config.EnableAuthentication {
		agentObjFromToken, err := GetClientNameFromToken(context,a.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		agentName=agentObjFromToken.Name
	}
	code,body:=a.agentService.Get(agentName)
	return context.JSON(code,body)
}

// NewAgentApi returns Agent type api
func NewAgentApi(agentService service.Agent, jwtService service.Jwt) api.Agent {
	return &agentApi{
		agentService: agentService,
		jwtService:         jwtService,
	}
}
