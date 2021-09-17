package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration
const (
	ServerHost = "server.host"
	ServerPort = "server.port"

	RedirectionURL = "redirection.url"
	CustomerIDType = "customer.id.type"
	CustomerID     = "customer.id"
)

func init() {
	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")
	_ = viper.BindEnv(RedirectionURL, "REDIRECTION_URL")
	_ = viper.BindEnv(CustomerIDType, "CUSTOMER_ID_TYPE")
	_ = viper.BindEnv(CustomerID, "CUSTOMER_ID")

	// defaults
	viper.SetDefault(ServerHost, "0.0.0.0")
	viper.SetDefault(ServerPort, "8081")
	viper.SetDefault(CustomerID, "awsTestID")
	viper.SetDefault(CustomerIDType, "STATIC")

	viper.SetDefault(RedirectionURL, "http://localhost/v1/aws/marketplace")
}
