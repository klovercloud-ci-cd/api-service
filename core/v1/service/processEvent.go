package service

type ProcessEvent interface {
	Store(events interface{}) (httpCode int, error error)
}
