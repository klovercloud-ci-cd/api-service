package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type podService struct {
	httpPublisher service.HttpClient
}

func (p podService) Get(action, companyId, agent, ownerReference, processId string, option v1.ResourceQueryOption) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.LighthouseQueryServerUrl+"/pods"+"?action="+action+"&companyId="+companyId+"&agent="+agent+"&owner-reference="+ownerReference+"&processId="+processId+"&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit+"&sort="+option.AscendingSort, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (p podService) GetByID(id, agent, processId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.LighthouseQueryServerUrl+"/pods/"+id+"?agent="+agent+"&processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

// NewNodeService returns Pod type service
func NewPodService(publisher service.HttpClient) service.Pod {
	return podService{
		httpPublisher: publisher,
	}
}
