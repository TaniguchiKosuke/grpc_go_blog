package server

import "grpc_go_blog/services/user/handler"

type API struct {
	UserHandler *handler.UserHandler
}

func NewAPI(userHandler *handler.UserHandler) *API {
	return &API{
		UserHandler: userHandler,
	}
}