package services

import (
	"context"
	"fmt"
	pb "wpc/user-service/internal/grpc"
	"wpc/user-service/internal/repositories"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	r repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) *UserService {
	return &UserService{r: r}
}

func (s *UserService) GetUserByID(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.r.GetUserByID(req.Id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return &pb.UserResponse{
		Id:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
	}, nil
}

func (s *UserService) LoginAttempt(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.r.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return &pb.LoginResponse{
		Id:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
