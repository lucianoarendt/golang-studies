package routes

import (
	"generic-list-crud-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	controllers := controllers.Start()
	//Tasks
	app.Post("/api/task/create", controllers.List.Create)
	app.Get("/api/task/read", controllers.List.Read)
	app.Put("/api/task/update", controllers.List.Update)
	app.Delete("/api/task/delete", controllers.List.Delete)
}
