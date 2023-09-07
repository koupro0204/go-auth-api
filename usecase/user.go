package usecase

import (
	"go-auth-api/domain/entity"
	"go-auth-api/domain/repository"
)

// interfaceから呼び出せるように
type User interface {
	InsertUser(email string, password string) (*entity.User, error)
}

type user struct {
	userRepository repository.User
}

// Userデータに対するusecaseを生成(依存関係の注入用)
func NewUserUseCase(ur repository.User) User {
	return &user{
		userRepository: ur,
	}
}
func (u user) InsertUser(email string, password string) (*entity.User, error) {
	//domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	// 本来ならここでpasswordのハッシュ化を行う。

	// データベースにユーザデータを登録する
	user, err := u.userRepository.InsertUser(&entity.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
