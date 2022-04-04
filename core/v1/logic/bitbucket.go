package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/api-service/config"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
)

type bitbucketService struct {
	httpPublisher service.HttpClient
}

//this function is responsible for forwarding the request to integration-manager
func (b bitbucketService) GetBranches(repoName, userName, repositoryId string) (httpCode int, body interface{}) {
	var response interface{}
	header := make(map[string]string)
	header["token"] = config.Token
	code, res, err := b.httpPublisher.Get(config.KlovercloudIntegrationMangerUrl+"bitbuckets?repoName="+repoName+"&userName="+userName+"&repoId="+repositoryId, header)
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
