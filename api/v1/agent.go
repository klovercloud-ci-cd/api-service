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
// @Router /api/v1/agents/{agent} [GET]
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

func (a agentApi) GetByName(context echo.Context) error {
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
	agent := context.Param("agent")
	code, data := a.agentService.GetByName(agent, companyId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Agents Query Failed", "Operation Failed")
}

// Get... Get K8sObjs Api
// @Summary Get K8sObjs api
// @Description Api for getting all K8sObjs short info by agent name and process id
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=v1.k8sObjsInfo{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/k8sobjs [GET]
func (a agentApi) GetK8sObjs(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetK8sObjs(agent, processId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

// Get... Get Pods by DaemonSet Api
// @Summary Get Pods by DaemonSet api
// @Description Api for getting all K8SPods by agent name, process id and DaemonSet uid
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "Certificate ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/daemonSets/{daemonSetId}/pods [GET]
func (a agentApi) GetPodsByDaemonSet(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	daemonSetId := context.Param("daemonSetId")
	if daemonSetId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Daemon Set ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByDaemonSet(agent, processId, daemonSetId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

// Get... Get Pods by Deployment Api
// @Summary Get Pods by Deployment api
// @Description Api for getting all K8SPods by agent name, process id and Deployment uid
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "Deployment ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/deployments/{deploymentId}/pods [GET]
func (a agentApi) GetPodsByDeployment(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	deploymentId := context.Param("deploymentId")
	if deploymentId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Deployment ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByDeployment(agent, processId, deploymentId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

// Get... Get Pods by ReplicaSet Api
// @Summary Get Pods by ReplicaSet api
// @Description Api for getting all K8SPods by agent name, process id and ReplicaSet uid
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "ReplicaSet ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/replicaSets/{replicaSetId}/pods [GET]
func (a agentApi) GetPodsByReplicaSet(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	replicaSetId := context.Param("replicaSetId")
	if replicaSetId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Replica Set ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByReplicaSet(agent, processId, replicaSetId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

// Get... Get Pods by StatefulSet Api
// @Summary Get Pods by StatefulSet api
// @Description Api for getting all K8SPods by agent name, process id and StatefulSet uid
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "StatefulSet ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/statefulSets/{statefulSetId}/pods [GET]
func (a agentApi) GetPodsByStatefulSet(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	statefulSetId := context.Param("statefulSetId")
	if statefulSetId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Stateful Set ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByStatefulSet(agent, processId, statefulSetId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
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
