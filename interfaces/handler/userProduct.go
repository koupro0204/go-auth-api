package handler

import (
	myerror "go-auth-api/error"
	"go-auth-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Userに対するHandlerのインターフェース
// main.goで使えるようにするため
type UserProduct interface {
	HandleUserProductCreate(c echo.Context) error
	HandleAuth(c echo.Context) error
}

type userProduct struct {
	userProductUseCase usecase.UserProduct
}

// Userデータに関するHandlerを生成
func NewUserProductHandler(up usecase.UserProduct) UserProduct {
	return &userProduct{
		userProductUseCase: up,
	}
}

type userProductRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Number   string `json:"number"`
}

// HandleUserProductCreate UserProduct作成処理
func (up userProduct) HandleUserProductCreate(c echo.Context) error {
	req := &userProductRequest{}
	if err := c.Bind(req); err != nil {
		return myerror.New(myerror.ErrorHttp, "create Bind error")
	}
	// emailのバリデーションを行う。
	if err := isValidEmail(req.Email); err != nil {
		return myerror.New(myerror.ErrorValidationEmail, "Invalid email")
	}
	// データベースにproductデータを登録する
	_, err := up.userProductUseCase.InsertUserProduct(req.Email, req.Password, req.Name, req.Number)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

type authResponse struct {
	Auth bool `json:"auth"`
}

// HandleAuth 認証処理
func (up userProduct) HandleAuth(c echo.Context) error {
	req := &userProductRequest{}
	if err := c.Bind(req); err != nil {
		return myerror.New(myerror.ErrorHttp, "create Bind error")
	}
	// emailのバリデーションを行う。
	if err := isValidEmail(req.Email); err != nil {
		return myerror.New(myerror.ErrorValidationEmail, "Invalid email")
	}
	// 認証処理
	if err := up.userProductUseCase.SelectUserProduct(req.Email, req.Password, req.Name, req.Number); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &authResponse{Auth: true})
}
