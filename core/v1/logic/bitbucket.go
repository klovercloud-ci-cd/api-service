package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"net/http"
)

type bitbucketService struct {
	httpPublisher service.HttpClient
}

func (b bitbucketService) EnableWebhook(companyId, repoId, userName, repoName string) (httpCode int, err error) {
	header := make(map[string]string)
	header["token"] = config.Token
	code, err := b.httpPublisher.Put(config.KlovercloudIntegrationMangerUrl+"/bitbuckets/webhook?repoName="+repoName+"&userName="+userName+"&repoId="+repoId+"&companyId="+companyId, header, nil)
	if err != nil {
		return code, err
	}
	return code, nil
}

func (b bitbucketService) DisableWebhook(companyId, repoId, userName, repoName, webhookId string) (httpCode int, err error) {
	header := make(map[string]string)
	header["token"] = config.Token
	code, err := b.httpPublisher.Delete(config.KlovercloudIntegrationMangerUrl+"/bitbuckets/webhook?repoName="+repoName+"&userName="+userName+"&repoId="+repoId+"&companyId="+companyId+"&webhookId="+webhookId, header)
	if err != nil {
		return code, err
	}
	return code, nil
}

func (b bitbucketService) GetCommitByBranch(repoName, userName, repoId, branch, companyId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, res, err := b.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/bitbuckets/branches?repoName="+repoName+"&userName="+userName+"&repoId="+repoId+"&companyId="+companyId+"&branch="+branch, header)
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
func (b bitbucketService) GetBranches(repoName, userName, repositoryId, companyId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, res, err := b.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"/bitbuckets/branches?repoName="+repoName+"&userName="+userName+"&repoId="+repositoryId+"&companyId="+companyId, header)
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
func (b bitbucketService) ListenEvent(payload interface{}, companyId string) error {
	marshal, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return marshalErr
	}
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"

	_, err := b.httpPublisher.Post(config.KlovercloudIntegrationMangerUrl+"/bitbuckets?companyId="+companyId, header, marshal)
	if err != nil {
		return err
	}

	return nil
}

// NewBitbucketService returns bitbucket type service
func NewBitbucketService(httpPublisher service.HttpClient) service.Bitbucket {
	return &bitbucketService{
		httpPublisher: httpPublisher,
	}
}
