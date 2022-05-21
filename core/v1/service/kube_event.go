package service

// KubeEvent k8s Event operations.
type KubeEvent interface {
	Store(events interface{}) (httpCode int, error error)
}
