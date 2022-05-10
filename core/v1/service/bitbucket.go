package service

type Bitbucket interface {
	ListenEvent(payload interface{}, companyId string) error
	GetBranches(repoName, userName, repositoryId, companyId string) (httpCode int, body interface{})
	GetCommitByBranch(repoName, userName, repoId, branch, companyId string) (httpCode int, body interface{})
}
