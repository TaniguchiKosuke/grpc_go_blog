package usecase

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"grpc_go_blog/domain/entity"
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
	userUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed to genereate hashed password")
		return nil, err
	}

	user := &entity.User{
		ID:       userUUID.String(),
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := uu.userRepository.CreateUser(user)
	if err != nil {
		log.Println("failed to create user")
		return nil, err
	}

	userOutput := &output.SignupUserResponse{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}

	return userOutput, nil
}

func (uu *UserUsecase) Login() {

}
