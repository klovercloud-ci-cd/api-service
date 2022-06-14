package service

// Agent operations
type Agent interface {
	Get(companyId string) (httpCode int, body interface{})
	GetTerminalByName(name string) (httpCode int, body interface{})
	Store(agent interface{}, name string) (httpCode int, error error)
}
