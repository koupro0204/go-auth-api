package main

import (
	"flag"
	"log"

	"go-auth-api/infrastructure/database"
	"go-auth-api/infrastructure/persistence"
	"go-auth-api/interfaces/handler"
	"go-auth-api/middleware"
	"go-auth-api/usecase"

	"github.com/labstack/echo/v4"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}
func main() {
	// データベース接続
	db, err := database.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	// User依存関係を注入
	userPersistence := persistence.NewUserPersistence(db)
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	e := echo.New()
	// 作成したerrorをミドルウェアに設定
	e.HTTPErrorHandler = middleware.ErrorHandler
	// /* ===== URLマッピングを行う ===== */
	e.POST("/user/create", userHandler.HandleUserCreate)

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server. %+v", err)
	}
}
