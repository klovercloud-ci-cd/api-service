package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type roleBindingService struct {
	httpPublisher service.HttpClient
}

func (r roleBindingService) Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := r.httpPublisher.Get(config.LighthouseQueryServerUrl+"/role-bindings"+"?agent="+agent+"&owner-reference="+ownerReference+"&processId="+processId+"&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit+"&sort="+option.AscendingSort, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (r roleBindingService) GetByID(id, agent, processId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := r.httpPublisher.Get(config.LighthouseQueryServerUrl+"/role-bindings/"+id+"?agent="+agent+"&processId="+processId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

// NewRoleBindingService returns RoleBinding type service
func NewRoleBindingService(publisher service.HttpClient) service.RoleBinding {
	return roleBindingService{
		httpPublisher: publisher,
	}
}
