package router

import (
	"ccg/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/ccg", handlers.GetCCG)
	e.POST("/ccg", handlers.AddCCG)
	e.PUT("/ccg", handlers.EditCCG)

	// e.GET("/ghoul", handlers.GetGhoul)
	// e.POST("/ghoul", handlers.AddGhoul)
}
