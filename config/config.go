package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

// ServerPort refers to server port.
var ServerPort string

// KlovercloudIntegrationMangerUrl refers to integration-managers url.
var KlovercloudIntegrationMangerUrl string

// KlovercloudEventStoreUrl refers to event-bank url.
var KlovercloudEventStoreUrl string

// KlovercloudEventStoreWebSocketUrl refers to event-stores socket url.
var KlovercloudEventStoreWebSocketUrl string

// PrivateKey refers to PrivateKey of EventStoreToken.
var PrivateKey string

// PublicKey refers to publickey of EventStoreToken.
var PublicKey string

// EnableAuthentication refers if service to service authentication is enabled.
var EnableAuthentication bool

// Token refers to jwt token for service to service communication.
var Token string

// InitEnvironmentVariables initializes environment variables
func InitEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR:", err.Error(), ", reading from env")

	}
	ServerPort = os.Getenv("SERVER_PORT")
	KlovercloudIntegrationMangerUrl = os.Getenv("KLOVERCLOUD_CI_INTEGRATION_MANAGER_URL")
	KlovercloudEventStoreUrl = os.Getenv("KLOVERCLOUD_CI_EVENT_STORE")
	PrivateKey = os.Getenv("PRIVATE_KEY")
	PublicKey = os.Getenv("PUBLIC_KEY")
	KlovercloudEventStoreWebSocketUrl = os.Getenv("KLOVERCLOUD_CI_EVENT_STORE_WS")

	if os.Getenv("ENABLE_AUTHENTICATION") == "" {
		EnableAuthentication = false
	} else {
		if strings.ToLower(os.Getenv("ENABLE_AUTHENTICATION")) == "true" {
			EnableAuthentication = true
		} else {
			EnableAuthentication = false
		}
	}
	Token = os.Getenv("TOKEN")
}
