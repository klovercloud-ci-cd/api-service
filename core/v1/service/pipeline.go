package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// Pipeline Pipeline operations.
type Pipeline interface {
	Get(companyId, repositoryId, url, revision, action, from, to string) (httpCode int, body interface{})
	GetByProcessId(companyId, processId, action string, option v1.Pagination) (httpCode int, body interface{})
	ReadEventsByCompanyId(c chan map[string]interface{}, companyId, userId string)
	ReadEventsByCompanyIdAndUserIdAndTime(c chan map[string]interface{}, companyId, userId, from string)
	Create(companyId, repositoryId, url string, payload interface{}) (httpCode int, body interface{})
	Update(companyId, repositoryId, url string, payload interface{}) (httpCode int, body interface{})
}
