package user

import (
	"context"
	"github.com/baransonmez/coff.app/app/web/grpc/user/pb"
	"github.com/baransonmez/coff.app/business/core/user"
	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedUserServer
	UserService *user.Service
}

func (s *Server) IsValidUser(ctx context.Context, r *pb.IsValidUserRequest) (*pb.IsValidResponse, error) {
	id, err := uuid.Parse(r.GetUuid())
	if err != nil {
		return &pb.IsValidResponse{IsValid: false}, err
	}
	_, err = s.UserService.GetUser(ctx, id)
	if err != nil {
		return &pb.IsValidResponse{IsValid: false}, err
	}
	return &pb.IsValidResponse{IsValid: true}, nil
}
