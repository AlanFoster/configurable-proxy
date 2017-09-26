package main

import (
	"log"
	"net/http"

	"github.com/alanfoster/configurable-proxy/proxy"
)

func main() {
	log.Println("Started application")

	http.Handle("/", proxy.Handler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
