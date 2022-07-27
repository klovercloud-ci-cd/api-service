package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// KubeObject k8s Object operations.
type KubeObject interface {
	Get(action, companyId, object, agent, ownerReference, processId string, option v1.ResourceQueryOption) (httpCode int, body interface{})
	GetByID(object, id, agent, processId string) (httpCode int, body interface{})
}
