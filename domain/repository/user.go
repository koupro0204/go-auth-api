package repository

import (
	"go-auth-api/domain/entity"
)

type User interface {
	InsertUser(user *entity.User) (*entity.User, error)
	SelectUserByEmail(email string) (*entity.User, error)
}
