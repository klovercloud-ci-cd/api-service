package service

type Github interface {
	ListenEvent(payload interface{}, companyId string) error
	GetBranches(url, repositoryId, companyId string) (httpCode int, body interface{})
	GetCommitByBranch(username, repositoryName, branch, companyId, repoId string) (httpCode int, body interface{})
	EnableWebhook(companyId, repoId, userName, repoName string) (httpCode int, err error)
	DisableWebhook(companyId, repoId, userName, repoName, webhookId string) (httpCode int, err error)
}
