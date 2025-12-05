package main

import (
	"github.com/Dawit0/examAuth/internal/delivery/handler"
	"github.com/Dawit0/examAuth/internal/infrastructure/database"
	repo "github.com/Dawit0/examAuth/internal/infrastructure/repository/userRepo"
	"github.com/Dawit0/examAuth/internal/infrastructure/route"
	"github.com/Dawit0/examAuth/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.DBconnection()

	userRepos := repo.NewUserRepo(db)
	resetUserRepo := repo.NewResetUserRepo(db)

	usecase := service.NewUserService(userRepos)
	resetUserUseCase := service.NewResetUserService(resetUserRepo, service.NewTestMailer("123456"))

	userHandler := handler.NewUserHandler(usecase)
	resetHandler := handler.NewForgetPasswordHandler(resetUserUseCase)

	routes := gin.Default()

	route.UserRoute(userHandler, routes)
	route.ResetRoute(resetHandler, routes)

	routes.Run(":8080")

}
