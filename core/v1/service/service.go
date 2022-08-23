package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// Service k8s Object operations.
type Service interface {
	Get(agent, ownerReference, processId string, option v1.ResourceQueryOption) (httpCode int, body interface{})
	GetByID(id, agent, processId string) (httpCode int, body interface{})
}