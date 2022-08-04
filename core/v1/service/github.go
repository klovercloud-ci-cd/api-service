package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

type Github interface {
	ListenEvent(payload interface{}, companyId, appId string) error
	GetBranches(url, repositoryId, companyId string) (httpCode int, body interface{})
	GetCommitByBranch(url, repoId, branch, companyId string, option v1.Pagination) (httpCode int, body interface{})
	EnableWebhook(companyId, repoId, url string) (httpCode int, err error)
	DisableWebhook(companyId, repoId, url, webhookId string) (httpCode int, err error)
}
