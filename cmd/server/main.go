package main

import (
	"github.com/matthewyuh246/socallogin/internal/controller"
	"github.com/matthewyuh246/socallogin/internal/repository"
	"github.com/matthewyuh246/socallogin/internal/router"
	"github.com/matthewyuh246/socallogin/internal/usecase"
	db "github.com/matthewyuh246/socallogin/pkg/database"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRespository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
