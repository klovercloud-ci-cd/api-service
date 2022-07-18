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

// Get... Get Pods by Certificate Api
// @Summary Get Pods by Certificate api
// @Description Api for getting all K8SPods by agent name, process id and Certificate id
// @Tags Agent
// @Produce json
// @Param agent path string true "Agent Name"
// @Param certificateId path string true "Certificate ID"
// @Param processId query string true "Process ID"
// @Success 200 {object} common.ResponseDTO{data=[]v1.K8sPod{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/agents/{agent}/certificates/{certificateId}/pods [GET]
func (a agentApi) GetPodsByCertificate(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	certificateId := context.Param("certificateId")
	if certificateId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Certificate ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByCertificate(agent, processId, certificateId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByClusterRole(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	clusterRoleId := context.Param("clusterRoleId")
	if clusterRoleId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Cluster Role ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByClusterRole(agent, processId, clusterRoleId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByClusterRoleBinding(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	clusterRoleBindingId := context.Param("clusterRoleBindingId")
	if clusterRoleBindingId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Cluster Role Binding ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByClusterRoleBinding(agent, processId, clusterRoleBindingId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByConfigMap(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	configMapId := context.Param("configMapId")
	if configMapId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Config Map ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByConfigMap(agent, processId, configMapId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

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

func (a agentApi) GetPodsByIngress(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	ingressId := context.Param("ingressId")
	if ingressId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Ingress ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByIngress(agent, processId, ingressId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByNamespace(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	namespaceId := context.Param("namespaceId")
	if namespaceId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Namespace ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByNamespace(agent, processId, namespaceId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByNetworkPolicy(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	networkPolicyId := context.Param("networkPolicyId")
	if networkPolicyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Network Policy ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByNetworkPolicy(agent, processId, networkPolicyId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByNode(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	nodeId := context.Param("nodeId")
	if nodeId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Node ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByNode(agent, processId, nodeId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByPV(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	pvId := context.Param("pvId")
	if pvId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Persistent Volume ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByPV(agent, processId, pvId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByPVC(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	pvcId := context.Param("pvcId")
	if pvcId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Persistent Volume ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByPVC(agent, processId, pvcId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

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

func (a agentApi) GetPodsByRole(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	roleId := context.Param("roleId")
	if roleId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Role ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByRole(agent, processId, roleId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByRoleBinding(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	roleBindingId := context.Param("roleBindingId")
	if roleBindingId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Role Binding ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByRoleBinding(agent, processId, roleBindingId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsBySecret(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	secretId := context.Param("secretId")
	if secretId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Secret ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsBySecret(agent, processId, secretId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByService(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	serviceId := context.Param("serviceId")
	if serviceId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Service ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByService(agent, processId, serviceId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

func (a agentApi) GetPodsByServiceAccount(context echo.Context) error {
	agent := context.Param("agent")
	if agent == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Agent name is not found", "Operation Failed")
	}
	processId := context.QueryParam("processId")
	if processId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Process ID is not found", "Operation Failed")
	}
	serviceAccountId := context.Param("serviceAccountId")
	if serviceAccountId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Service Account ID is not found", "Operation Failed")
	}
	code, data := a.agentService.GetPodsByServiceAccount(agent, processId, serviceAccountId)
	if code == 200 {
		return context.JSON(200, data)
	}
	return common.GenerateErrorResponse(context, "Query Failed", "Operation Failed")
}

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
