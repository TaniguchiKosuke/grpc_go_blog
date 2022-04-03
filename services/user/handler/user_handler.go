package handler

import (
	"context"
	"log"

	"grpc_go_blog/services/user/handler/proto"
	"grpc_go_blog/services/user/usecase"
	"grpc_go_blog/services/user/usecase/input"
)

type UserHandler struct {
	proto.UnimplementedUserServiceServer

	userUsecase usecase.UserUsecaseIF
}

func NewUserHandler(userUsecase usecase.UserUsecaseIF) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (uh *UserHandler) Signup(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	user := &input.SignupUserRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := uh.userUsecase.Signup(user)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &proto.SignupResponse{
		Id:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}, nil
}
