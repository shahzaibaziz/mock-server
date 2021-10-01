package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"github.com/shahzaibaziz/mock-server/config"
	"github.com/spf13/viper"
)

type AwsCustomer struct {
	CustomerIdentifier string
	ProductCode        string
}

// CustomerResolveHandler  returns  aws customer id and product id
func CustomerResolveHandler(w http.ResponseWriter, r *http.Request) {
	var data AwsCustomer
	if viper.GetString(config.CustomerIDType) == "STATIC" {
		data = AwsCustomer{ProductCode: "657mykebcauj9r2yp9zkb16nm", CustomerIdentifier: viper.GetString(config.CustomerID)}
	} else {
		data = AwsCustomer{CustomerIdentifier:uuid.New().String(), ProductCode:"657mykebcauj9r2yp9zkb16nm"}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
