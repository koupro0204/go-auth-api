package persistence

import (
	"database/sql"
	"go-auth-api/domain/entity"
	"go-auth-api/domain/repository"
	myerror "go-auth-api/error"
)

type userPersistence struct {
	DB *sql.DB
}

func NewUserPersistence(DB *sql.DB) repository.User {
	return &userPersistence{DB: DB}
}

func (u userPersistence) InsertUser(user *entity.User) (*entity.User, error) {
	if _, err := u.DB.Exec(
		"INSERT INTO User (email, password) values (?, ?)",
		user.Email,
		user.Password,
	); err != nil {
		return nil, myerror.New(myerror.ErrorDB, "Failed user insert")
	}
	return user, nil
}
