package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/labstack/echo/v4"
)

type kubeEvent struct {
	kubeEventService service.KubeEvent
	jwtService       service.Jwt
}

// Save... Save k8s event
// @Summary Save k8s event
// @Description Stores k8s event
// @Tags KubeEvent
// @Accept json
// @Produce json
// @Param data body interface{} true "KubeEvent Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/kube_events [POST]
func (k kubeEvent) Save(context echo.Context) error {
	var formData interface{}
	if err := context.Bind(&formData); err != nil {
		return err
	}
	if config.EnableAuthentication {
		_, err := GetClientNameFromBearerToken(context)
		if err != nil {
			return context.JSON(401, "Unauthorized user!")
		}
	}
	_, err := k.kubeEventService.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, err, err.Error())
	}
	return common.GenerateSuccessResponse(context, "[SUCCESS]: k8s events saved successfully", nil, "k8s events saved!")
}

// NewKubeEventApi returns KubeEvent type api
func NewKubeEventApi(kubeEventService service.KubeEvent, jwtService service.Jwt) api.KubeEvent {
	return &kubeEvent{
		kubeEventService: kubeEventService,
		jwtService:       jwtService,
	}
}
