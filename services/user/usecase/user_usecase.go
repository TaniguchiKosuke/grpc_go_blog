package usecase

import (
	"grpc_go_blog/domain/repository"
	"grpc_go_blog/services/user/usecase/input"
	"grpc_go_blog/services/user/usecase/output"
)

type UserUsecaseIF interface {
	Signup(*input.SignupUserRequest) (*output.SignupUserResponse, error)
	Login()
}

type UserUsecase struct {
	userRepository repository.UserRepositoryIF
}

func NewUserUsecase(userRepository repository.UserRepositoryIF) UserUsecaseIF {
	return &UserUsecase{userRepository: userRepository}
}

func (uu *UserUsecase) Signup(userInput *input.SignupUserRequest) (*output.SignupUserResponse, error) {
	return nil, nil
}

func (uu *UserUsecase) Login() {

}
