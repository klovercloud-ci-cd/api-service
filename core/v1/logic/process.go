package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
)

type processService struct {
	httpPublisher service.HttpClient
}

func (p processService) GetByCompanyIdAndRepositoryIdAndAppName(companyId string, repositoryId string, appId string, option v1.ProcessQueryOption) (httpCode int, body interface{}) {
	var response interface{}

	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/processes?"+"companyId="+companyId+"&repositoryId="+repositoryId+"&appId="+appId+"&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit+"&step="+option.Step, header)

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
