package main

import (
	"log"
	"net/http"

	"github.com/ilhamdcp/dcp-go-microservice/common"
)

var (
	httpPort = common.GetEnvString("HTTP_PORT", ":8081")
)

func main() {
	mux := http.NewServeMux()
	httpHandler := NewHttpHandler()
	httpHandler.registerRoutes(mux)

	log.Printf("Start http server at port %s", httpPort)

	err := http.ListenAndServe(httpPort, mux)
	if err != nil {
		log.Fatal("Failed to serve HTTP on port: {}")
	}
}
