package routing

import (
	"Eight/src/internal/delivery/http"
	"Eight/src/internal/middleware"
	"Eight/src/internal/service"
	"github.com/labstack/echo/v4"
)

func SetupUserRoute(e *echo.Echo, userService *service.UserService) {
	userHandler := http.NewUserHandler(userService)

	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.ErrorHandlerMiddleware)
	apiGroup.POST("/users", userHandler.Create)
	apiGroup.GET("/users", userHandler.GetAll)
	apiGroup.GET("/users/:id", userHandler.GetById)
	apiGroup.PUT("/users", userHandler.Update)
	apiGroup.DELETE("/users", userHandler.Delete)
}
