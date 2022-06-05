package service

import v1 "github.com/klovercloud-ci-cd/api-service/core/v1"

// KubeObject k8s Object operations.
type KubeObject interface {
	Get(object, agent, ownerReference, processId string, option v1.ResourceQueryOption) (httpCode int, body interface{})
}
