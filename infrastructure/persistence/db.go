package persistence

import (
	"gapi/domain/repository"
)

type Repositories struct {
	User repository.UserRepository
}

func NewRepositories() (*Repositories, error) {
	return &Repositories{
		User: NewUserRepository(),
	}, nil
}
