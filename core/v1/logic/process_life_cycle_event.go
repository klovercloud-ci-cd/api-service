package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"log"
	"net/http"
)

type processLifeCycleEventService struct {
	httpPublisher service.HttpClient
}

func (p processLifeCycleEventService) Store(events interface{}) (httpCode int, error error) {
	marshal, marshalErr := json.Marshal(events)
	if marshalErr != nil {
		return http.StatusBadRequest, marshalErr
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"
	code, err := p.httpPublisher.Post(config.KlovercloudEventStoreUrl+"/process_life_cycle_events", header, marshal)
	if err != nil {
		return code, err
	}
	return code, nil
}

func (p processLifeCycleEventService) PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count , stepType string) (httpCode int, body interface{}) {
	response := make(map[string]interface{})
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/process_life_cycle_events"+"?step_type="+stepType+"&count="+count, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}
	return code, response
}

func (p processLifeCycleEventService) PullPausedAndAutoTriggerEnabledResourcesByAgentName(count , agent string) (httpCode int, body interface{}) {
	response := make(map[string]interface{})
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"
	code, b, err := p.httpPublisher.Get(config.KlovercloudEventStoreUrl+"/process_life_cycle_events"+"?agent="+agent+"&count="+count, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, nil
	}
	return code, response
}

// NewProcessLifeCycleEventService returns ProcessLifeCycleEvent type service
func NewProcessLifeCycleEventService(publisher service.HttpClient) service.ProcessLifeCycleEvent {
	return processLifeCycleEventService{
		httpPublisher: publisher,
	}
}
