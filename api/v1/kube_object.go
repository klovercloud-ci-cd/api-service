package v1

import (
	"github.com/klovercloud-ci-cd/api-service/api/common"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/api"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"github.com/klovercloud-ci-cd/api-service/enums"
	"github.com/labstack/echo/v4"
)

type kubeObject struct {
	kubeObjectService service.KubeObject
	jwtService        service.Jwt
}

// Get... Get Api
// @Summary Get api
// @Description Api for getting all kube objects by object name, agent name, owner reference and process id
// @Tags KubeObject
// @Produce json
// @Param owner-reference query string false "Owner Reference"
// @Param processId query string true "Process Id"
// @Param agent query string true "Agent Name"
// @Param page query int64 false "Page Number"
// @Param limit query int64 false "Limit"
// @Param sort query bool false "Sort By Created Time"
// @Success 200 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/kube_objects [GET]
func (k kubeObject) Get(context echo.Context) error {
	if config.EnableAuthentication {
		userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, k.jwtService)
		if err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
		if err := checkAuthority(userResourcePermission, string(enums.PROCESS), "", string(enums.READ)); err != nil {
			return common.GenerateUnauthorizedResponse(context, err, err.Error())
		}
	}
	object := context.QueryParam("object")
	object = reformatObjectName(object)
	agent := context.QueryParam("agent")
	ownerReference := context.QueryParam("owner-reference")
	processId := context.QueryParam("processId")
	option := getK8sObjectQueryOption(context)
	code, data := k.kubeObjectService.Get(object, agent, ownerReference, processId, option)
	if code == 200 {
		return context.JSON(code, data)
	}
	return common.GenerateErrorResponse(context, "k8s Object Query Failed", "Operation Failed")
}

func reformatObjectName(object string) string {
	switch object {
	case "certificate":
		return "certificates"
	case "cluster-role":
		return "cluster-roles"
	case "cluster-role-binding":
		return "cluster-role-bindings"
	case "config-map":
		return "config-maps"
	case "daemon-set":
		return "daemon-sets"
	case "deployment":
		return "deployments"
	case "ingress":
		return "ingresses"
	case "namespace":
		return "namespaces"
	case "network-policy":
		return "network-policies"
	case "node":
		return "nodes"
	case "pod":
		return "pods"
	case "persistent-volume":
		return "persistent-volumes"
	case "persistent-volume-claim":
		return "persistent-volume-claims"
	case "replica-set":
		return "replica-sets"
	case "role":
		return "roles"
	case "role-binding":
		return "role-bindings"
	case "secret":
		return "secrets"
	case "service":
		return "services"
	case "service-account":
		return "service-accounts"
	case "stateful-set":
		return "stateful-sets"
	}
	return ""
}

//this function is for set all query param
func getK8sObjectQueryOption(context echo.Context) v1.ResourceQueryOption {
	option := v1.ResourceQueryOption{}
	option.Pagination.Page = context.QueryParam("page")
	option.Pagination.Limit = context.QueryParam("limit")
	option.AscendingSort = context.QueryParam("sort")
	return option
}

// NewKubeObjectApi returns KubeObject type api
func NewKubeObjectApi(kubeObjectService service.KubeObject, jwtService service.Jwt) api.KubeObject {
	return &kubeObject{
		kubeObjectService: kubeObjectService,
		jwtService:        jwtService,
	}
}
