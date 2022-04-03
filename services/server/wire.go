//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"grpc_go_blog/config/db"
	userhandler "grpc_go_blog/services/user/handler"
	usersqlite "grpc_go_blog/services/user/infra/sqlite"
	userusecase "grpc_go_blog/services/user/usecase"
)

var userSuperSet = wire.NewSet(
	userhandler.NewUserHandler,
	userusecase.NewUserUsecase,
	usersqlite.NewUserRepository,
)

func InjectAPIServer() (*API, error) {
	wire.Build(
		NewAPI,
		db.NewDB,
		userSuperSet,
	)

	return nil, nil
}