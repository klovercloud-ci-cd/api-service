package service

// WebsocketClient WebsocketClient operations.
type WebsocketClient interface {
	Get(c chan map[string]interface{}, url string, header map[string][]string)
}
