package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type kubeObjectService struct {
	httpPublisher service.HttpClient
}

func (k kubeObjectService) Get(action, companyId, object, agent, ownerReference, processId string, option v1.ResourceQueryOption) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := k.httpPublisher.Get(config.LighthouseQueryServerUrl+"/"+object+"?action="+action+"&companyId="+companyId+"&agent="+agent+"&owner-reference="+ownerReference+"&processId="+processId+"&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit+"&sort="+option.AscendingSort, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

// NewKubeObjectService returns KubeObject type service
func NewKubeObjectService(publisher service.HttpClient) service.KubeObject {
	return kubeObjectService{
		httpPublisher: publisher,
	}
}
