package error

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type errorcode int

const (
	// 良い感じに定義する
	ErrorUnExpected errorcode = 1000
	ErrorNotFound   errorcode = 404
	ErrorHttp       errorcode = 3000
	ErrorDB         errorcode = 4000
	ErrorQuery      errorcode = 5000
	ErrorValidation errorcode = 6000
)

type Exception struct {
	Code           int
	Message        string
	File           string
	Line           int
	About          string
	HttpStatusCode int
}

func New(code errorcode, message string) *Exception {
	_, file, line, _ := runtime.Caller(1)
	// カレントディレクトリを取得
	currentDir, _ := filepath.Abs(".")
	// 相対パスを計算
	relativePath, err := filepath.Rel(currentDir, file)
	if err != nil {
		return &Exception{
			Code:    int(ErrorUnExpected),
			Message: fmt.Sprintf("Failed to calculate relative path: %s", err),
		}
	}

	res := &Exception{
		Code:    int(code),
		Message: message,
		File:    relativePath,
		Line:    line,
	}
	switch code {
	case ErrorHttp:
		res.About = "http"
		res.HttpStatusCode = 500
	case ErrorDB:
		res.About = "db"
		res.HttpStatusCode = 500
	}
	return res
}

func (e *Exception) Error() string {
	return fmt.Sprintf("%s:%d %s: %s", e.File, e.Line, e.About, e.Message)
}
