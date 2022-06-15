package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type ProcessEvent struct {
	processEvent service.ProcessEvent
	jwtService   service.Jwt
}

// Get ... Get Proccess Event By Company Id And Process Id
// @Summary Get Proccess Event By Company Id And Process Id
// @Description Get Proccess Event By Company Id And Process Id
// @Tags ProcessEvent
// @Accept json
// @Produce json
// @Param scope query string false "scope [notification]"
// @Param companyId query string true "Company Id"
// @Param processId query string false "Process Id when scope is notification [Optional]"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/processes_events [GET]
func (p ProcessEvent) Get(context echo.Context) error {
	var companyId string
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		companyId = userResourcePermission.Metadata.CompanyId
	}
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR] No companyId is provided", "Operation failed")
	}
	processId := context.QueryParam("processId")
	scope := context.QueryParam("scope")
	option := getProcessQueryOption(context)
	code, data := p.processEvent.Get(companyId, processId, scope, option)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "Process event query failed", "Operation Failed")
}

// Save... Save Pipeline process event
// @Summary Save Pipeline process event
// @Description Stores Pipeline process event
// @Tags ProcessEvent
// @Accept json
// @Produce json
// @Param data body interface{} true "PipelineProcessEvent Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/processes_events [POST]
func (p ProcessEvent) Save(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}

	if config.EnableAuthentication {
		_, err := GetClientNameFromToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
	}
	code, err := p.processEvent.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if code == 200 {
		return common.GenerateSuccessResponse(context, nil, nil, "Pipeline Process Event Saved Successfully")
	}
	return common.GenerateErrorResponse(context, nil, err.Error())
}

func NewProcessEvent(processEvent service.ProcessEvent, jwtService service.Jwt) api.ProcessEvent {
	return &ProcessEvent{
		processEvent: processEvent,
		jwtService:   jwtService,
	}
}
