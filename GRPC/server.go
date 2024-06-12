package main

import (
	"context"
	"log"
	"net"

	pb "grpc-test/multiplex" // Import the generated package

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type server struct {
	pb.UnimplementedMultiplexServiceServer
}

func (s *server) Process(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	// Extract client IP address from the context
	clientIP := getClientIP(ctx)
	log.Printf("Received request from client IP: %s", clientIP)

	// log.Printf("Received request: %v", req)
	log.Printf("Processing request ID: %s with data: %s", req.Id, req.Data)
	result := "Processed: " + req.Data
	return &pb.Response{Id: req.Id, Result: result}, nil
}

// Function to extract client IP address from the context
func getClientIP(ctx context.Context) string {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		return "unknown"
	}
	return peer.Addr.String()
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterMultiplexServiceServer(grpcServer, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
