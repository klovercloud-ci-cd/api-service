package service

type Github interface {
	ListenEvent(payload interface{}, companyId string) error
	GetBranches(repoName, userName, repositoryId string) (httpCode int, body interface{})
}
