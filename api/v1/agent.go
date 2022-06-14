package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type agentApi struct {
	agentService service.Agent
	jwtService   service.Jwt
}

// Get.. Get Agents
// @Summary Get Agents by company id
// @Description Get Agents by company id
// @Tags Agent
// @Produce json
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/agents [GET]
func (a agentApi) Get(context echo.Context) error {
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, a.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Company Id is not found.", "Operation failed.")
	}
	code, data := a.agentService.Get(companyId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Agents Query Failed", "Operation Failed")
}

// Save... Save Agents terminal information
// @Summary  Save Agents terminal information
// @Description Save Agents terminal information
// @Tags Agent
// @Accept json
// @Produce json
// @Param data body object true "Agents Terminal Data"
// @Param name query string false "agent name"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/agents [POST]
func (a agentApi) Save(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	agentName := context.QueryParam("name")
	if config.EnableAuthentication {
		agentObjFromToken, err := GetClientNameFromToken(context, a.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		agentName = agentObjFromToken.Name
	}
	code, err := a.agentService.Store(formData, agentName)
	if err != nil {
		return common.GenerateErrorResponse(context, err, err.Error())
	}
	return context.JSON(code, err)
}

// Get.. Get Agents terminal info by agent name
// @Summary Get Agents terminal info by agent name
// @Description Get Agents terminal info by agent name
// @Tags Agent
// @Produce json
// @Param name path string true "agent name"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/agents/{name} [GET]
func (a agentApi) GetTerminalByName(context echo.Context) error {
	agentName := context.Param("name")
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, a.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
	}
	code, body := a.agentService.GetTerminalByName(agentName)
	return context.JSON(code, body)
}

// NewAgentApi returns Agent type api
func NewAgentApi(agentService service.Agent, jwtService service.Jwt) api.Agent {
	return &agentApi{
		agentService: agentService,
		jwtService:   jwtService,
	}
}
