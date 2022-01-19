package v1

import (
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
)

type processLifeCycleEventApi struct {
	processLifeCycleEventService service.ProcessLifeCycleEvent
}

// Pull... Pull Steps
// @Summary Pull Steps
// @Description Pulls auto trigger enabled steps
// @Tags ProcessLifeCycle
// @Produce json
// @Param agent path string true "Agen name"
// @Param count path int64 true "Pull size"
// @Param step_type path string false "Step type [BUILD, DEPLOY]"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/process_life_cycle_events [GET]
func (p processLifeCycleEventApi) Pull(context echo.Context) error {
	agentName := context.QueryParam("agent")
	count := context.QueryParam("count")
	steptype := context.QueryParam("step_type")
	if steptype != "" {
		return context.JSON(p.processLifeCycleEventService.PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count, steptype))
	}
	return context.JSON(p.processLifeCycleEventService.PullPausedAndAutoTriggerEnabledResourcesByAgentName(count, agentName))
}

// Save ... Save process lifecycle event
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
	return context.JSON(p.processLifeCycleEventService.Store(formData))
}

// NewProcessLifeCycleEventApi returns ProcessLifeCycleEvent type api
func NewProcessLifeCycleEventApi(processLifeCycleEventService service.ProcessLifeCycleEvent) api.ProcessLifeCycleEvent {
	return &processLifeCycleEventApi{
		processLifeCycleEventService: processLifeCycleEventService,
	}
}
