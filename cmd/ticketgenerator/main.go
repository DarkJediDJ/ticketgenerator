package main

import (
	"log"
	"net"

	"github.com/darkjedidj/ticketgenerator/api/server"
	pb "github.com/darkjedidj/ticketgenerator/internal/proto"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to lsiten: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTicketGeneratorServer(s, &server.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to lsiten: %v", err)
	}
}
