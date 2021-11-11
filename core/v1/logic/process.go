package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
)

type processService struct {
	httpPublisher service.HttpClient
}

func (p processService) GetByCompanyIdAndRepositoryIdAndAppName(companyId string, repositoryId string, appId string) (httpCode int, body interface{}) {
	var response interface{}

	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/processes?"+"companyId="+companyId+"&repositoryId="+repositoryId+"&appId="+appId, header)

	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(b, &response)
	if er != nil {
		return code, nil
	}
	return code, response
}

// NewProcessService returns Process type service
func NewProcessService(publisher service.HttpClient) service.Process {
	return processService{
		httpPublisher: publisher,
	}
}
