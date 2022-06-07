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

func (a agentService) Get(name string) (httpCode int, body interface{}) {
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
