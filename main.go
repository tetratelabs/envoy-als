package main

import (
	sink "envoy-als/pkg"
	"fmt"
	"log"
	"net"
	"os"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"google.golang.org/grpc"
)

var defaultPort = "5000"

func main() {

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}
	grpcServer := grpc.NewServer()
	v3.RegisterAccessLogServiceServer(grpcServer, sink.New())

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on :%s", port)
	grpcServer.Serve(l)
}
