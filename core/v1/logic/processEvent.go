package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type ProcessEvent struct {
	httpPublisher service.HttpClient
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
