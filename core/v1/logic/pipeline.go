package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type pipelineService struct {
	httpClient      service.HttpClient
	websocketClient service.WebsocketClient
}

func (p pipelineService) Create(companyId, repositoryId, url string, payload interface{}) (httpCode int, body interface{}) {
	marshal, err := json.Marshal(payload)
	if err != nil {
		return http.StatusBadRequest, err
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"

	code, err := p.httpClient.Post(config.KlovercloudIntegrationMangerUrl+"/pipelines?url="+url+"&companyId="+companyId+"&repositoryId="+repositoryId, header, marshal)
	return code, err
}

func (p pipelineService) Update(companyId, repositoryId, url string, payload interface{}) (httpCode int, body interface{}) {
	marshal, err := json.Marshal(payload)
	if err != nil {
		return http.StatusBadRequest, err
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"

	code, err := p.httpClient.Put(config.KlovercloudIntegrationMangerUrl+"/pipelines?url="+url+"&companyId="+companyId+"&repositoryId="+repositoryId, header, marshal)
	return code, err
}

func (p pipelineService) ReadEventsByCompanyId(c chan map[string]interface{}, companyId string) {
	p.websocketClient.Get(c, config.KlovercloudEventStoreWebSocketUrl+"/pipelines/ws?company_id="+companyId, http.Header{"token": {config.Token}})
}
func (p pipelineService) Get(companyId, repositoryId, url, revision, action string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token

	code, b, err := p.httpClient.Get(config.KlovercloudIntegrationMangerUrl+"/pipelines?action="+action+"&companyId="+companyId+"&repositoryId="+repositoryId+"&url="+url+"&revision="+revision, header)

	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (p pipelineService) GetByProcessId(processId, action string, option v1.Pagination) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token

	code, b, err := p.httpClient.Get(config.KlovercloudEventStoreUrl+"/pipelines/"+processId+"?action="+action+"&order=&page="+option.Page+"&limit="+option.Limit, header)

	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

// NewPipelineService returns Pipeline type service
func NewPipelineService(publisher service.HttpClient, websocketClient service.WebsocketClient) service.Pipeline {
	return pipelineService{
		httpClient:      publisher,
		websocketClient: websocketClient,
	}
}
