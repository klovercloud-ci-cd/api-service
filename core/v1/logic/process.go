package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type processService struct {
	httpPublisher service.HttpClient
}

func (p processService) GetProcessLifeCycleEventByProcessIdAndStepName(companyId,processId, step string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/processes/"+processId+"/process_life_cycle_events?step="+step+"&companyId="+companyId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (p processService) GetById(companyId, processId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/processes/"+processId+"?companyId="+companyId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (p processService) GetLogsByProcessIdAndStepAndFootmark(processId, step, footmark string, claims string, option v1.CompanyQueryOption) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/processes/"+processId+"/steps/"+step+"/footmarks/"+footmark+"/logs?claims="+claims+"&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (p processService) GetFootmarksByProcessIdAndStep(processId, step string) (httpCode int, body interface{}) {
	var response interface{}

	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/processes/"+processId+"/steps/"+step+"/footmarks", header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (p processService) Get(companyId, repositoryId, appId, commitId, operation, from, to string, option v1.ProcessQueryOption) (httpCode int, body interface{}) {
	var response interface{}

	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/processes?"+"companyId="+companyId+"&repositoryId="+repositoryId+"&appId="+appId+"&commitId="+commitId+"&operation="+operation+"&from="+from+"&to"+to+"&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit+"&step="+option.Step, header)

	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

// NewProcessService returns Process type service
func NewProcessService(publisher service.HttpClient) service.Process {
	return processService{
		httpPublisher: publisher,
	}
}
