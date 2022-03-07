package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// Process Process operations.
type Process interface {
	GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string, option v1.ProcessQueryOption) (httpCode int, body interface{})
}
