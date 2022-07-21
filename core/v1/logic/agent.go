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
