package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
)

type processLifeCycleEventApi struct {
	processLifeCycleEventService service.ProcessLifeCycleEvent
	jwtService                   service.Jwt
}

// Pull... Pull Steps
// @Summary Pull Steps
// @Description Pulls auto trigger enabled steps
// @Tags ProcessLifeCycle
// @Produce json
// @Param agent query string true "Agen name"
// @Param count query int64 true "Pull size"
// @Param step_type query string false "Step type [BUILD, DEPLOY]"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/process_life_cycle_events [GET]
func (p processLifeCycleEventApi) Pull(context echo.Context) error {
	agentName := context.QueryParam("agent")
	count := context.QueryParam("count")
	steptype := context.QueryParam("step_type")
	if config.EnableAuthentication {
		client, err := GetClientNameFromToken(context, p.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if steptype != "" {
			code, data := p.processLifeCycleEventService.PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count, steptype)
			if code == 200 {
				return context.JSON(200, data)
			}
			return common.GenerateErrorResponse(context, "Steps Query Failed", "Operation Failed")
		}
		code, data := p.processLifeCycleEventService.PullPausedAndAutoTriggerEnabledResourcesByAgentName(count, client.Name)
		if code == 200 {
			return context.JSON(200, data)
		}
		return common.GenerateErrorResponse(context, "Steps Query Failed", "Operation Failed")
	} else {
		if steptype != "" {
			code, data := p.processLifeCycleEventService.PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count, steptype)
			if code == 200 {
				return context.JSON(200, data)
			}
			return common.GenerateErrorResponse(context, "Steps Query Failed", "Operation Failed")
		}
		code, data := p.processLifeCycleEventService.PullPausedAndAutoTriggerEnabledResourcesByAgentName(count, agentName)
		if code == 200 {
			return context.JSON(code, data)
		}
		return common.GenerateErrorResponse(context, "Steps Query Failed", "Operation Failed")
	}
}

// Save... Save process lifecycle event
// @Summary Save process lifecycle event
// @Description Stores process lifecycle event
// @Tags ProcessLifeCycle
// @Accept json
// @Produce json
// @Param data body interface{} true "ProcessLifeCycleEventList Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/process_life_cycle_events [POST]
func (p processLifeCycleEventApi) Save(context echo.Context) error {
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
	code, err := p.processLifeCycleEventService.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if code == 200 {
		return common.GenerateSuccessResponse(context, nil, nil, "Process Life Cycle Event Saved Successfully")
	}
	return common.GenerateErrorResponse(context, nil, err.Error())
}

// NewProcessLifeCycleEventApi returns ProcessLifeCycleEvent type api
func NewProcessLifeCycleEventApi(processLifeCycleEventService service.ProcessLifeCycleEvent, jwtService service.Jwt) api.ProcessLifeCycleEvent {
	return &processLifeCycleEventApi{
		processLifeCycleEventService: processLifeCycleEventService,
		jwtService:                   jwtService,
	}
}
