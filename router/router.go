package router

import (
	"os"

	"github.com/DarcoProgramador/auth-go/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var c *fiber.App

func Init(userHandler *user.Handler) {
	c = fiber.New()

	c.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST",
		AllowHeaders:     "Content-Type",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	c.Use(logger.New())

	api := c.Group("/api")

	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	api.Get("/user", userHandler.Me)
	api.Get("/logout", userHandler.Logout)

}

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	c.Listen(":" + port)
}
