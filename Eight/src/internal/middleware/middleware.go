package middleware

import (
	"Eight/src/internal/customError"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			log.Println("Error occurred:", err)
			if appErr, ok := err.(*customError.AppError); ok {
				return c.JSON(appErr.Code, appErr)
			}
			return c.JSON(http.StatusInternalServerError, customError.New(http.StatusInternalServerError, "Internal Server Error"))
		}
		return nil
	}
}
