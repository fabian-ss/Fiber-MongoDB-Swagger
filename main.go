package main

import (
	"github.com/fabian-ss/Fiber-MongoDB-Swagger/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	docs "github.com/fabian-ss/Fiber-MongoDB-Swagger/docs"
	r "github.com/fabian-ss/Fiber-MongoDB-Swagger/routes"
)

// @contact.name   API Support - My Github
// @contact.url    https://github.com/fabian-ss
// @contact.email  randommusicd@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	docs.SwaggerInfo.Title = "API con Fiber, Swagger y MongoDB"
	docs.SwaggerInfo.Description = "CRUD de usuarios usando mongoDb, la base de datos se llama golangAPI"
	docs.SwaggerInfo.Host = "http://localhost" + config.FiberPort()
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := fiber.New()

	config.ConnectDB()

	app.Use(cors.New())

	r.DocRoute(app)

	r.UserRoute(app)

	r.NotFoundRoute(app)

	app.Listen(config.FiberPort())
}
