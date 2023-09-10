package handler

import (
	myerror "go-auth-api/error"
	"go-auth-api/usecase"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v4"
)

// Userに対するHandlerのインターフェース
// main.goで使えるようにするため
type User interface {
	HandleUserCreate(c echo.Context) error
}

type user struct {
	userUseCase usecase.User
}

// Userデータに関するHandlerを生成
func NewUserHandler(uu usecase.User) User {
	return &user{
		userUseCase: uu,
	}
}

type userCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// HandleUserCreate ユーザ作成処理
func (u user) HandleUserCreate(c echo.Context) error {
	req := &userCreateRequest{}
	if err := c.Bind(req); err != nil {
		return myerror.New(myerror.ErrorHttp, "create Bind error")
	}
	// emailのバリデーションを行う。
	if err := isValidEmail(req.Email); err != nil {
		return myerror.New(myerror.ErrorValidationEmail, "Invalid email")
	}
	// データベースにユーザデータを登録する
	_, err := u.userUseCase.InsertUser(req.Email, req.Password)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func isValidEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
