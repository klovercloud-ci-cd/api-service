package service

type LogEvent interface {
	Store(log interface{}) (httpCode int, error error)
}
