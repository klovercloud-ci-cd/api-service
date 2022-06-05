package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
)

type LogEvent struct {
	logEventService service.LogEvent
	jwtService      service.Jwt
}

// Save... Save log
// @Summary Save log
// @Description Stores logs
// @Tags Log
// @Accept json
// @Produce json
// @Param data body interface{} true "LogEvent Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/logs [POST]
func (l LogEvent) Save(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}

	if config.EnableAuthentication {
		_, err := GetClientNameFromBearerToken(context)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
	}

	code, err := l.logEventService.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if code == 200 {
		return common.GenerateSuccessResponse(context, nil, nil, "Log Saved Successfully")
	}
	return common.GenerateErrorResponse(context, nil, err.Error())
}

func NewLogEvent(logEventService service.LogEvent, jwtService service.Jwt) api.LogEvent {
	return &LogEvent{
		logEventService: logEventService,
		jwtService:      jwtService,
	}
}
