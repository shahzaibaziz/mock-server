package handlers

import (
	"fmt"
	"github.com/shahzaibaziz/mock-server/config"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// CustomerOnBoardingHandler  redirect browser to VPC+ marketplace url
func CustomerOnBoardingHandler(w http.ResponseWriter, r *http.Request) {
	target := fmt.Sprintf("%s?aws_token=testtoken123", viper.GetString(config.RedirectionURL))
	log.Printf("redirect to: %s", target)
	http.Redirect(w, r, target, 302)
}
