package repository

import (
	"go-auth-api/domain/entity"
)

type User interface {
	InsertUser(user *entity.User) (*entity.User, error)
}
