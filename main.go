package main

import (
	"ccg/router"
	"ccg/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	storage.InitDB()
	e := echo.New()
	router.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
