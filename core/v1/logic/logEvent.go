package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type LogEvent struct {
	httpPublisher service.HttpClient
}

func (l LogEvent) Store(log interface{}) (httpCode int, error error) {
	marshal, marshalErr := json.Marshal(log)
	if marshalErr != nil {
		return http.StatusBadRequest, marshalErr
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"
	code, err := l.httpPublisher.Post(config.KlovercloudEventStoreUrl+"/logs", header, marshal)
	if err != nil {
		return code, err
	}
	return code, nil
}

func NewLogEventService(httpPublisher service.HttpClient) service.LogEvent {
	return &LogEvent{
		httpPublisher: httpPublisher,
	}
}
