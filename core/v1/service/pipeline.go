package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// Pipeline Pipeline operations.
type Pipeline interface {
	Get(companyId, repositoryId, url, revision, action string) (httpCode int, body interface{})
	GetByProcessId(processId, action string, option v1.Pagination) (httpCode int, body interface{})
	ReadEventsByCompanyId(c chan map[string]interface{}, companyId string)
}
