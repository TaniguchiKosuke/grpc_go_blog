package handler

import "grpc_go_blog/services/user/handler/proto"

type UserHandler struct {
	proto.UnimplementedUserServiceServer

}
