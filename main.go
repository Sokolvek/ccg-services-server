package main

import (
	"ccg/repository"
	"ccg/router"
	"ccg/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	// feddis.InitDB()
	storage.InitDB()
	e := echo.New()
	router.InitRoutes(e)
	repository.InitRepo()
	e.Logger.Fatal(e.Start(":8080"))
}
