package middleware

import (
	"errors"
	myerror "go-auth-api/error"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// LogError はエラー情報をログに出力します。
func LogError(e *myerror.Exception) {
	log.Printf("Error Code: %d, Message: %s, File: %s, Line: %d, About: %s\n",
		e.Code, e.Message, e.File, e.Line, e.About)
}

func ErrorHandler(err error, c echo.Context) {
	var e *myerror.Exception

	// EchoのHTTPErrorを処理
	if he, ok := err.(*echo.HTTPError); ok {
		if he.Code == http.StatusNotFound {
			e = myerror.New(myerror.ErrorNotFound, "Resource not found")
			LogError(e)
			c.JSON(http.StatusNotFound, ErrorResponse{
				Code:    e.Code,
				Message: e.Error(),
			})
			return
		}
	}

	// エラータイプを確認
	if !errors.As(err, &e) {
		e = myerror.New(myerror.ErrorUnExpected, "Unexpected error")
	}

	// エラーログを出力
	LogError(e)

	// JSONレスポンスを返す
	if jsonErr := c.JSON(e.HttpStatusCode, ErrorResponse{
		Code:    e.Code,
		Message: e.Error(),
	}); jsonErr != nil {
		log.Fatalf("Failed to send JSON response: %v", jsonErr)
	}
}
