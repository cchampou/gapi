package persistence

import (
	"gapi/domain/entity"
	"gapi/domain/repository"
)

type UserRepo struct{}

func NewUserRepository() *UserRepo {
	return &UserRepo{}
}

var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) GetUser(id string) (*entity.User, error) {
	return &entity.User{ID: id, Nickname: "k_talpa"}, nil
}
