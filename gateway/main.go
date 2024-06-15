package main

import "net/http"

const (
	httpAddress = "8080"
)

func main() {
	mux := http.NewServeMux()
	httpHandler := NewHttpHandler()
	httpHandler.registerRoutes(mux)

	err := http.ListenAndServe(httpAddress, httpHandler)
}
