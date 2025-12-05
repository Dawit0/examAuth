package main

import (
	"github.com/Dawit0/examAuth/internal/delivery/handler"
	"github.com/Dawit0/examAuth/internal/infrastructure/database"
	repo "github.com/Dawit0/examAuth/internal/infrastructure/repository/userRepo"
	"github.com/Dawit0/examAuth/internal/infrastructure/route"
	"github.com/Dawit0/examAuth/internal/pkg/logger"
	"github.com/Dawit0/examAuth/internal/server/middleware"
	"github.com/Dawit0/examAuth/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger()
	db := database.DBconnection()

	logger.Logger.Info("Database connected successfully")

	userRepos := repo.NewUserRepo(db)
	resetUserRepo := repo.NewResetUserRepo(db)

	usecase := service.NewUserService(userRepos)
	resetUserUseCase := service.NewResetUserService(resetUserRepo, service.NewGmailMailer("workenhdawit@gmail.com", "vlvs ygcl odpe gzee"))

	userHandler := handler.NewUserHandler(usecase)
	resetHandler := handler.NewForgetPasswordHandler(resetUserUseCase)

	routes := gin.Default()

	routes.Use(
		middleware.LoggingMiddleware(logger.Logger),
		middleware.RecoveryMiddleware(logger.Logger),
	)

	route.UserRoute(userHandler, routes)
	route.ResetRoute(resetHandler, routes)

	logger.Logger.Info("Server started on :8080")

	routes.Run(":8080")

}
