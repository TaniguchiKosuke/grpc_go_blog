package handler

import (
	"context"
	"grpc_go_blog/services/user/handler/proto"
	"grpc_go_blog/services/user/usecase"
)

type UserHandler struct {
	proto.UnimplementedUserServiceServer

	useUsecase usecase.UserUsecaseIF
}

func NewUserHandler(userUsecase usecase.UserUsecaseIF) *UserHandler {
	return &UserHandler{useUsecase: userUsecase}
}

func (uh *UserHandler) Signup(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	return nil, nil
}
