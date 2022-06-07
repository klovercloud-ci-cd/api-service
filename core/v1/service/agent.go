package service

// Agent operations
type Agent interface {
	Store(agent interface{}, name string) (httpCode int, error error)
	Get(name string) (httpCode int, body interface{})
}
