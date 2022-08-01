package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type namespaceService struct {
	httpPublisher service.HttpClient
}

func (n namespaceService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := n.httpPublisher.Get(config.LighthouseQueryServerUrl+"/namespaces"+"?agent="+agent+"&owner-reference="+ownerReference+"&processId="+processId+"&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit+"&sort="+option.AscendingSort, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (n namespaceService) GetByID(id, agent, processId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := n.httpPublisher.Get(config.LighthouseQueryServerUrl+"/namespaces/"+id+"?agent="+agent+"&processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

// NewNamespaceService returns Namespace type service
func NewNamespaceService(publisher service.HttpClient) service.Namespace {
	return namespaceService{
		httpPublisher: publisher,
	}
}
