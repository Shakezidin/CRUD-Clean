package di

import (
	"github.com/shakezidin/internal/database"
	"github.com/shakezidin/internal/routes"
	"github.com/shakezidin/internal/server"
	"github.com/shakezidin/internal/user/delivery"
	"github.com/shakezidin/internal/user/repository"
	"github.com/shakezidin/internal/user/usecase"
)

func Init() *server.ServerStruct {
	server := server.NewHTTPServer()
	db := database.ConnectDatabase()
	repo := repository.NewUserRepository(db)
	usecase := usecase.UseCase(repo)
	delivery := delivery.UserDelivery(usecase)
	userRoutes := routes.NewUserInit(server, delivery)
	userRoutes.UserRoutes()
	return server
}
