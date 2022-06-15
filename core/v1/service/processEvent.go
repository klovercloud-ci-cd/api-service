package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

type ProcessEvent interface {
	Get(companyId, processId, scope string, option v1.ProcessQueryOption) (httpCode int, data interface{})
	Store(events interface{}) (httpCode int, error error)
}
