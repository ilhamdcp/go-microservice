package main

import (
	"log"
	"net/http"

	"github.com/ilhamdcp/go-microservice/common"
	pb "github.com/ilhamdcp/go-microservice/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpPort            = common.GetEnvString("HTTP_PORT", ":8081")
	orderServiceAddress = "localhost:3000"
)

func main() {
	conn, err := grpc.NewClient(orderServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to initiate GRPC service, error: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	httpHandler := NewHttpHandler(client)
	httpHandler.registerRoutes(mux)

	log.Printf("Start http server at port %s", httpPort)

	err = http.ListenAndServe(httpPort, mux)
	if err != nil {
		log.Fatal("Failed to serve HTTP on port: {}")
	}
}
