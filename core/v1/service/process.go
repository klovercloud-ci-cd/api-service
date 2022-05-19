package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// Process Process operations.
type Process interface {
	Get(companyId, repositoryId, appId, commitId string, option v1.ProcessQueryOption) (httpCode int, body interface{})
	GetFootmarksByProcessIdAndStep(processId, step string) (httpCode int, body interface{})
	GetLogsByProcessIdAndStepAndFootmark(processId, step, footmark string, claims string, option v1.CompanyQueryOption) (httpCode int, body interface{})
}
