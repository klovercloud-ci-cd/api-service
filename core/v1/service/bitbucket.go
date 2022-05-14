package service

type Bitbucket interface {
	ListenEvent(payload interface{}, companyId string) error
	GetBranches(url, repositoryId, companyId string) (httpCode int, body interface{})
	GetCommitByBranch(url, repoId, branch, companyId string) (httpCode int, body interface{})
	EnableWebhook(companyId, repoId, url string) (httpCode int, err error)
	DisableWebhook(companyId, repoId, url, webhookId string) (httpCode int, err error)
}
