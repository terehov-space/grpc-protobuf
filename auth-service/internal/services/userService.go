package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
	pb "wpc/auth-service/internal/grpc"
)

type UserService struct {
	client pb.UserServiceClient
}

func NewUserService(c *grpc.ClientConn) *UserService {
	return &UserService{
		client: pb.NewUserServiceClient(c),
	}
}

func (s *UserService) GetUserByID(c *gin.Context) {
	reqId := c.Param("id")
	id, err := strconv.ParseInt(reqId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	resp, err := s.client.GetUserByID(context.Background(), &pb.UserRequest{
		Id: id,
	})
	if err != nil {
		log.Printf("Error when calling GetUserByID: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user"})
	}

	c.JSON(http.StatusOK, resp)
}

func (s *UserService) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := s.client.LoginAttempt(context.Background(), &pb.LoginRequest{
		Email: req.Email,
	})

	if err != nil {
		log.Printf("Error when calling Login: %v\n", err)
		c.JSON(http.StatusForbidden, gin.H{"error": "Request is forbidden"})
		return
	}

	// TODO: logic to check password
	if req.Password != "secret" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Request is forbidden"})
		return
	}

	// TODO: generate JWT token for frontend

	c.JSON(http.StatusOK, resp)
}
