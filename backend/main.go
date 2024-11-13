package main

import (
	"log"

	"github.com/posiposi/project/backend/controller"
	"github.com/posiposi/project/backend/db"
	"github.com/posiposi/project/backend/repository"
	"github.com/posiposi/project/backend/router"
	"github.com/posiposi/project/backend/usecase"
)

func main() {
	db := db.NewDB()
	log.Println("Successfully connected to database")
	// リポジトリ層コンストラクタ起動
	userRepository := repository.NewUserRepository(db)
	// ユースケース層コンストラクタ起動
	userUsecase := usecase.NewUserUsecase(userRepository)
	// コントローラ層コンストラクタ起動
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
