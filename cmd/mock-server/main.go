package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/shahzaibaziz/mock-server/config"
	"github.com/shahzaibaziz/mock-server/handlers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/aws-marketplace", handlers.CustomerOnBoardingHandler)
	r.HandleFunc("/customer-resolve", handlers.CustomerResolveHandler)

	log.Printf("server is running on 0.0.0.0:%s", "8081")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s",
		viper.GetString(config.ServerHost), viper.GetString(config.ServerPort)), r))
}
