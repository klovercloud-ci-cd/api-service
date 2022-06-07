package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
)

type ProcessEvent struct {
	processEvent service.ProcessEvent
	jwtService   service.Jwt
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
		_, err := GetClientNameFromToken(context,p.jwtService)
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
