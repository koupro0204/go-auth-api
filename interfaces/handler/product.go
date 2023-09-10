package handler

import (
	myerror "go-auth-api/error"
	"go-auth-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Userに対するHandlerのインターフェース
// main.goで使えるようにするため
type Product interface {
	HandleProductCreate(c echo.Context) error
}

type product struct {
	productUseCase usecase.Product
}

// Userデータに関するHandlerを生成
func NewProductHandler(up usecase.Product) Product {
	return &product{
		productUseCase: up,
	}
}

type productCreateRequest struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

// HandleProductCreate Product作成処理
func (p product) HandleProductCreate(c echo.Context) error {
	req := &productCreateRequest{}
	if err := c.Bind(req); err != nil {
		return myerror.New(myerror.ErrorHttp, "create Bind error")
	}

	// データベースにproductデータを登録する
	_, err := p.productUseCase.InsertProduct(req.Name, req.Number)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
