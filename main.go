package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "markplatts.org/scrud/protos"
	"markplatts.org/scrud/server"
)

const (
	port = 49000
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterScrudServer(grpcServer, &server.ScrudServer{})
	log.Printf("Server listening on port: %v", port)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}

}
