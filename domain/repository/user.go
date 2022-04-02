package repository

import "grpc_go_blog/domain/entity"

type UserRepositoryIF interface {
	CreateUser(*entity.User) (*entity.User, error)
	GetUserByEmail(string) (*entity.User, error)
}
