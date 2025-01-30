package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"wpc/user-service/internal/database"
	pb "wpc/user-service/internal/grpc"
	"wpc/user-service/internal/repositories"
	"wpc/user-service/internal/services"
)

func Serve() {
	db := database.InitDB()
	userRepository := repositories.UserRepository{DB: db}

	server := grpc.NewServer()
	userService := services.NewUserService(userRepository)

	pb.RegisterUserServiceServer(server, userService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Println("User service is running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
