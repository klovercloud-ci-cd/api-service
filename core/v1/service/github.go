package service

type Github interface {
	ListenEvent(payload interface{}, companyId string) error
	GetBranches(repoName, userName, repositoryId, companyId string) (httpCode int, body interface{})
	GetCommitByBranch(username, repositoryName, branch, companyId, repoId string) (httpCode int, body interface{})
}
