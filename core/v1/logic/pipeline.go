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

func (p pipelineService) ReadEventsByProcessId(c chan map[string]interface{}, processId string) {
	p.websocketClient.Get(c, config.KlovercloudEventStoreWebSocketUrl+"/pipelines/ws?processId="+processId, http.Header{"token": {config.Token}})
}

func (p pipelineService) GetByProcessId(processId string, option v1.CompanyQueryOption) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token

	code, b, err := p.httpClient.Get(config.KlovercloudEventStoreUrl+"/pipelines/"+processId+"?order=&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit, header)

	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(b, &response)
	if er != nil {
		return code, nil
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
