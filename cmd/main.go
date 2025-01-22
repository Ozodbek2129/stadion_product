package main

import (
	"log"
	"net"
	"stadion/config"
	pb "stadion/genproto/stadium"
	"stadion/pkg/logger"
	"stadion/service"
	"stadion/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	config := config.Load()
	listener, err := net.Listen("tcp", config.STADIUM_SERVICE)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	logs := logger.NewLogger()

	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	stadiumService := service.NewStadiumService(db, logs)

	server := grpc.NewServer()
	pb.RegisterStadiumServiceServer(server, stadiumService)

	log.Printf("Starting server at %s", config.STADIUM_SERVICE)

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
