package main

import (
	"log"
	"net/http"

	"github.com/mengdage/hello-go/apitest/handlers"
)

func main() {
	handlers.Routes()
	log.Println("listener: Started: Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
