package service

import (
	"context"
	"log"

	pb "account/api/rbac/v1"
)

type V1Service struct {
	pb.UnimplementedV1Server
}

func NewV1Service() *V1Service {
	return &V1Service{}
}

func (s *V1Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	log.Printf("Login in, %+v\n", req)
	return &pb.LoginReply{Message: "login success !" + req.UserName}, nil
}
func (s *V1Service) Me(ctx context.Context, req *pb.MeRequest) (*pb.MeReply, error) {
	log.Printf("Me in, %+v\n", req)
	return &pb.MeReply{}, nil
}
