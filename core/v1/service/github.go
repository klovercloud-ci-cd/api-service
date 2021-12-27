package service

type Github interface {
	ListenEvent(payload interface{}, companyId string) error
}
