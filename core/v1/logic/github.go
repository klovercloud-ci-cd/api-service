package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type githubService struct {
	httpPublisher service.HttpClient
}

func (g githubService) EnableWebhook(companyId, repoId, url string) (httpCode int, err error) {
	header := make(map[string]string)
	header["token"] = config.Token
	code, err := g.httpPublisher.Put(config.KlovercloudIntegrationMangerUrl+"/githubs/webhook?url="+url+"&repoId="+repoId+"&companyId="+companyId, header, nil)
	if err != nil {
		return code, err
	}
	return code, nil
}

func (g githubService) DisableWebhook(companyId, repoId, url, webhookId string) (httpCode int, err error) {
	header := make(map[string]string)
	header["token"] = config.Token
	code, err := g.httpPublisher.Delete(config.KlovercloudIntegrationMangerUrl+"/githubs/webhook?url="+url+"&repoId="+repoId+"&companyId="+companyId+"&webhookId="+webhookId, header)
	if err != nil {
		return code, err
	}
	return code, nil
}

func (g githubService) GetCommitByBranch(url, repoId, branch, companyId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, res, err := g.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/githubs/commits?url="+url+"&repoId="+repoId+"&companyId="+companyId+"&branch="+branch, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return code, response
}

//this function is responsible for forwarding the request to integration-manager
func (g githubService) GetBranches(url, repositoryId, companyId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, res, err := g.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/githubs/branches?url="+url+"&repoId="+repositoryId+"&companyId="+companyId, header)
	if err != nil {
		return code, err
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return http.StatusBadRequest, err
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
