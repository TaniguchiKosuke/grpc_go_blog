package sqlite

import (
	"grpc_go_blog/config/db"
	"grpc_go_blog/domain/entity"
	"grpc_go_blog/domain/repository"
)

type UserRepository struct {
	*db.DB
}

func NewUserRepository(db *db.DB) repository.UserRepositoryIF {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if err := ur.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	return nil, nil
}
