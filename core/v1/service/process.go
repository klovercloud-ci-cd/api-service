package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// Process Process operations.
type Process interface {
	Get(companyId, repositoryId, appId, commitId, operation, from, to string, option v1.ProcessQueryOption) (httpCode int, body interface{})
	GetById(companyId, processId string) (httpCode int, body interface{})
	GetProcessLifeCycleEventByProcessIdAndStepName(companyId, processId, step string) (httpCode int, body interface{})
	GetFootmarksByProcessIdAndStep(processId, companyId, step string) (httpCode int, body interface{})
	GetLogsByProcessIdAndStepAndFootmark(companyId, processId, step, footmark string, claims string, option v1.CompanyQueryOption) (httpCode int, body interface{})
}
