package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
)

type githubService struct {
	httpPublisher service.HttpClient
}

//this function is responsible for forwarding the request to integration-manager
func (g githubService) GetBranches(repoName, userName, repositoryId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, res, err := g.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"githubs?repoName="+repoName+"&userName="+userName+"&repoId="+repositoryId, header)
	if err != nil {
		return code, nil
	}
	er := json.Unmarshal(res, &response)
	if er != nil {
		return code, nil
	}
	return code, response
}

//this function is responsible for forwarding the request to integration-manager
func (g githubService) ListenEvent(payload interface{}, companyId string) error {
	marshal, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return marshalErr
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"

	_, err := g.httpPublisher.Post(config.KlovercloudIntegrationMangerUrl+"/githubs?companyId="+companyId, header, marshal)
	if err != nil {
		return err
	}

	return nil
}

// NewGithubService returns github type service
func NewGithubService(httpPublisher service.HttpClient) service.Github {
	return &githubService{
		httpPublisher: httpPublisher,
	}
}
