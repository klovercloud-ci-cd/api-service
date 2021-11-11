package service

// Process Process operations.
type Process interface {
	GetByCompanyIdAndRepositoryIdAndAppName(companyId string, repositoryId string, appId string) (httpCode int, body interface{})
}
