package repository

import "gapi/domain/entity"

type UserRepository interface {
	GetUser(string) (*entity.User, error)
}
