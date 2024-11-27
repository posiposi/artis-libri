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
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewBookRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	bookUsecase := usecase.NewBookUsecase(taskRepository)
	userController := controller.NewUserController(userUsecase)
	bookController := controller.NewBookController(bookUsecase)
	e := router.NewRouter(userController, bookController)
	e.Logger.Fatal(e.Start(":8081"))
}
