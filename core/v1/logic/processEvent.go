package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	v1 "github.com/klovercloud-ci-cd/api-service/core/v1"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type ProcessEvent struct {
	httpPublisher service.HttpClient
}

func (p ProcessEvent) Get(companyId, processId, scope string, option v1.ProcessQueryOption) (httpCode int, data interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/processes_events?companyId="+companyId+"&processId="+processId+"&scope="+scope+"&page="+option.Pagination.Page+"&limit="+option.Pagination.Limit, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

func (p ProcessEvent) Store(events interface{}) (httpCode int, error error) {
	marshal, marshalErr := json.Marshal(events)
	if marshalErr != nil {
		return http.StatusBadRequest, marshalErr
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"
	code, err := p.httpPublisher.Post(config.KlovercloudEventStoreUrl+"/processes_events", header, marshal)
	if err != nil {
		return code, err
	}
	return code, nil
}

func NewProcessEvent(httpPublisher service.HttpClient) service.ProcessEvent {
	return &ProcessEvent{
		httpPublisher: httpPublisher,
	}
}
