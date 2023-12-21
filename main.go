package main

import (
	"github.com/DarcoProgramador/auth-go/db"
	"github.com/DarcoProgramador/auth-go/internal/user"
	"github.com/DarcoProgramador/auth-go/router"
)

func main() {
	db := db.NewDB()
	repo := user.NewRepository(db)
	service := user.NewService(repo)
	userHandler := user.NewHandler(service)

	router.Init(userHandler)
	router.Start()
}
