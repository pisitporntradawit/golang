package main

import (
	"api/config"
	"api/module/login"
	"api/module/user"
	"api/route"
)

func main() {
	DB := config.Config()
	defer DB.Close()
	repo := user.NewRepository(DB)
	service := user.NewService(repo)
	controllor := user.NewControllor(service)

	loginRepo := login.NewRepository(DB)
	loginService := login.NewService(loginRepo)
	loginControllor := login.NewControllor(loginService)

	router := route.RouteAPI(controllor, loginControllor)
	router.Run(":30606")
}