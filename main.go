package main

import (
	"api/config"
	"api/module/user"
	"api/route"
)

func main() {
	DB := config.Config()
	defer DB.Close()
	repo := user.NewRepository(DB)
	service := user.NewService(repo)
	controllor := user.NewControllor(service)

	router := route.RouteAPI(controllor)
	router.Run(":30606")
}