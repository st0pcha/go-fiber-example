package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/st0pcha/go-fiber-example/internal/controllers"
)

func RegisterRoutes(app *fiber.App) {
	mainController(app)
	userController(app)
}

func mainController(app *fiber.App) {
	g := app.Group("/main")
	g.Get("/", controllers.GetHelloWorld)
}

func userController(app *fiber.App) {
	g := app.Group("/users")
	g.Get("/:id", controllers.GetUserByID)
}
