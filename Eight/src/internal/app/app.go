package app

import (
	"Eight/src/internal/configs"
	"Eight/src/internal/repository"
	"Eight/src/internal/routing"
	"Eight/src/internal/service"
	"Eight/src/pkg/database"
	"github.com/labstack/echo/v4"
)

func Run() {
	dbModel := configs.GetDbParams()
	db := database.NewClient(dbModel)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	e := echo.New()
	routing.SetupUserRoute(e, userService)

	e.Logger.Fatal(e.Start(":8080"))
}
