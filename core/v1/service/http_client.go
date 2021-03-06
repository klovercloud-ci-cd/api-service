package service

// HttpClient HttpClient operations.
type HttpClient interface {
	Get(url string, header map[string]string) (httpCode int, body []byte, err error)
	Post(url string, header map[string]string, body []byte) (httpCode int, err error)
	Put(url string, header map[string]string, body []byte) (httpCode int, err error)
	Delete(url string, header map[string]string) (httpCode int, err error)
	Patch(url string, header map[string]string, body []byte) (httpCode int, err error)
}
