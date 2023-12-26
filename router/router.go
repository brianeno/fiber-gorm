package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	_ "github.com/brianeno/projects-fiber/docs"
	projectRoutes "github.com/brianeno/projects-fiber/internals/routes/project"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/swagger/*", swagger.HandlerDefault)
	// Group api calls with param '/api'
	api := app.Group("/api", logger.New())

	// Setup project routes, can use same syntax to add routes for more models
	projectRoutes.SetupProjectRoutes(api)
}
