package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// Pipeline Pipeline operations.
type Pipeline interface {
	GetByProcessId(processId string, option v1.Pagination) (httpCode int, body interface{})
	ReadEventsByProcessId(c chan map[string]interface{}, processId string)
}
