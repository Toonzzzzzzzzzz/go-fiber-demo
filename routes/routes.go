package routes

import (
	"github.com/Toonzzzzzzzzzz/go-fiber-demo/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/test", handlers.Test)
	api.Get("/users", handlers.GetUsers)
	api.Post("/users/add", handlers.CreateUser)
	api.Delete("/users/:id", handlers.DeleteUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Get("/users/:id", handlers.GetUsersById)

}
