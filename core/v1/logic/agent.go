package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type agentService struct {
	httpClient service.HttpClient
}

func (a agentService) Get(companyId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents?companyId="+companyId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetK8sObjs(agent, processId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/k8sobjs?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByCertificate(agent, processId, certificateId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/certificates/"+certificateId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByClusterRole(agent, processId, clusterRoleId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/clusterRoles/"+clusterRoleId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByClusterRoleBinding(agent, processId, clusterRoleBindingId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/clusterRoleBindings/"+clusterRoleBindingId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByConfigMap(agent, processId, configMapId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/configMaps/"+configMapId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByDaemonSet(agent, processId, daemonSetId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/daemonSets/"+daemonSetId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByDeployment(agent, processId, deploymentId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/deployments/"+deploymentId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByIngress(agent, processId, ingressId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/ingresses/"+ingressId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByNamespace(agent, processId, namespaceId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/namespaces/"+namespaceId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByNetworkPolicy(agent, processId, networkPolicyId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/networkPolicies/"+networkPolicyId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByNode(agent, processId, nodeId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/nodes/"+nodeId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByPV(agent, processId, PVId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/persistentVolumes/"+PVId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByPVC(agent, processId, PVCId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/persistentVolumeClaims/"+PVCId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByReplicaSet(agent, processId, replicaSetId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/replicaSets/"+replicaSetId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByRole(agent, processId, roleId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/roles/"+roleId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByRoleBinding(agent, processId, roleBindingId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/roleBindings/"+roleBindingId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsBySecret(agent, processId, secretId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/secrets/"+secretId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByService(agent, processId, serviceId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/services/"+serviceId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByServiceAccount(agent, processId, serviceAccountId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/serviceAccounts/"+serviceAccountId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) GetPodsByStatefulSet(agent, processId, statefulSetId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := a.httpClient.Get(config.LighthouseQueryServerUrl+"/agents/"+agent+"/statefulSets/"+statefulSetId+"/pods?processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (a agentService) Store(agent interface{}, name string) (httpCode int, error error) {
	marshal, err := json.Marshal(agent)
	if err != nil {
		return http.StatusBadRequest, err
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"
	code, err := a.httpClient.Post(config.KlovercloudIntegrationMangerUrl+"/agents?name="+name, header, marshal)
	return code, err
}

func (a agentService) GetTerminalByName(name string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token

	code, b, err := a.httpClient.Get(config.KlovercloudIntegrationMangerUrl+"/agents/"+name, header)

	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

// NewAgentService returns agent type service
func NewAgentService(httpPublisher service.HttpClient) service.Agent {
	return &agentService{
		httpClient: httpPublisher,
	}
}
