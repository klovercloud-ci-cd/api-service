package logic

import (
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
)

type kubeEventService struct {
	httpPublisher service.HttpClient
}

func (k kubeEventService) Store(events interface{}) (httpCode int, error error) {
	//marshal, marshalErr := json.Marshal(events)
	//if marshalErr != nil {
	//	return http.StatusBadRequest, marshalErr
	//}
	//header := make(map[string]string)
	//header["token"] = config.Token
	//header["Content-Type"] = "application/json"
	//code, err := k.httpPublisher.Post(config.KlovercloudEventStoreUrl+"/kube_events", header, marshal)
	//if err != nil {
	//	return code, err
	//}
	//return code, nil
	panic("implement me!")
}

// NewKubeEventEventService returns KubeEvent type service
func NewKubeEventEventService(publisher service.HttpClient) service.KubeEvent {
	return kubeEventService{
		httpPublisher: publisher,
	}
}