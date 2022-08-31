package routes

import (
	ctrl "github.com/fabian-ss/Fiber-MongoDB-Swagger/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	routes := app.Group("/apiUsers")
	routes.Get("/user/:userId", ctrl.GetAUser)
	routes.Get("/users", ctrl.GetAllUsers)
	routes.Post("/user", ctrl.CreateUser)
	routes.Put("/user/:userId", ctrl.EditAUser)
	routes.Delete("/user/:userId", ctrl.DeleteAUser)

}
