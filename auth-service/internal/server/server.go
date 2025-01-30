package server

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"wpc/auth-service/internal/services"
)

func Serve() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	userService := services.NewUserService(conn)

	r := gin.Default()

	r.GET("/users/:id", userService.GetUserByID)
	r.POST("/auth/login", userService.Login)

	r.Run(":8080")
}
