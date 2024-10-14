package rest

import (
	"audit-log/src/interfaces/rest/controller"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes sets up the HTTP routes
func RegisterRoutes(app *echo.Echo) {
	log := app.Group("/audit-log")
	log.POST("/create", controller.CreateLog)
}
