package logic

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/klovercloud-ci-cd/api-service/core/v1/service"
	"log"
)

type websocketClientService struct {
}

func (w websocketClientService) Get(c chan map[string]interface{}, url string, header map[string][]string) {
	channel := make(map[string]interface{})
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Println("dial:", err)
		return
	}
	defer conn.Close()
	for {
		_, m, err := conn.ReadMessage()
		if err != nil {
			log.Println("[ERROR]: Failed to read:", err)
			conn.Close()
			return
		}
		err = json.Unmarshal(m, &channel)
		if err != nil {
			log.Println("[ERROR] Failed to Unmarshal data from socket:", err.Error())
			conn.Close()
			return
		}
		c <- channel
	}

}

// NewWebsocketClientService returns WebsocketClient type service
func NewWebsocketClientService() service.WebsocketClient {
	return &websocketClientService{}
}
