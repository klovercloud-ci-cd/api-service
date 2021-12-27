package service

type Bitbucket interface {
	ListenEvent(payload interface{}, companyId string) error
}
