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
		return nil, myerror.New(myerror.ErrorInsert, "Failed user insert")
	}
	return user, nil
}

// SelectUserByEmail emailを条件にレコードを取得する
func (up userPersistence) SelectUserByEmail(email string) (*entity.User, error) {
	// passwordが正しいかはusecaseで行う
	row := up.DB.QueryRow("SELECT * FROM User WHERE email=?", email)
	return convertToUser(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToUser(row *sql.Row) (*entity.User, error) {
	var user entity.User
	if err := row.Scan(&user.UserID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, myerror.New(myerror.ErrorNotExistUser, "User is not exist")
		}
		return nil, myerror.New(myerror.ErrorDB, "Failed convert to User")
	}
	return &user, nil
}
