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

	usecase := service.NewUserService(userRepos)

	handler := handler.NewUserHandler(usecase)

	routes := gin.Default()

	route.UserRoute(handler, routes)

	routes.Run(":9090")

}
