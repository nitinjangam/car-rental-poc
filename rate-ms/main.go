package main

import (
	"log"
	"net"

	"github.com/nitinjangam/car-rental-poc/definitions-store/ratepb"
	"github.com/nitinjangam/car-rental-poc/rate-ms/rateserver"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("error while listening on port 50051: %v", err.Error())
	}
	log.Println("rate server started listening on port 50052")
	rateServer := grpc.NewServer()
	ratepb.RegisterRoutingServiceServer(rateServer, &rateserver.Server{})
	if err := rateServer.Serve(lis); err != nil {
		log.Fatalf("error while rateServer Serve: %v", err.Error())
	}
}
