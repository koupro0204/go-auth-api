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
	// Product依存関係を注入
	productPersistence := persistence.NewProductPersistence(db)
	productUseCase := usecase.NewProductUseCase(productPersistence)
	productHandler := handler.NewProductHandler(productUseCase)
	// UserProduct依存関係を注入
	userProductPersistence := persistence.NewUserProductPersistence(db)
	userProductUseCase := usecase.NewUserProductUseCase(userPersistence, productPersistence, userProductPersistence)
	userProductHandler := handler.NewUserProductHandler(userProductUseCase)

	e := echo.New()
	// 作成したerrorをミドルウェアに設定
	e.HTTPErrorHandler = middleware.ErrorHandler
	// /* ===== URLマッピングを行う ===== */
	e.POST("/user/create", userHandler.HandleUserCreate)
	e.POST("/user/products", userProductHandler.HandleUserProductCreate)
	e.POST("/user/auth", userProductHandler.HandleAuth)

	e.POST("/products/create", productHandler.HandleProductCreate)

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server. %+v", err)
	}
}
