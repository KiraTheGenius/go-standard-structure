package main

import (
	"net/http"
	"standard-go-project/internal/services"

	"standard-go-project/internal/platform/logging"
)

func main() {
	logging.Logger.Info("Starting application")
	http.HandleFunc("/hello/{name}", services.Greet)
	http.ListenAndServe(":8081", nil)
}
