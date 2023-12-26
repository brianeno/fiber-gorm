package projectsRoutes

import (
	projectHandler "github.com/brianeno/projects-fiber/internals/handlers/project"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectRoutes(router fiber.Router) {
	project := router.Group("/project")
	// Create a Project
	project.Post("/", projectHandler.CreateProject)
	// Read all Projects
	project.Get("/", projectHandler.GetProjects)
	// // Read one Project
	project.Get("/:projectId", projectHandler.GetProject)
	// // Update one Project
	project.Put("/:projectId", projectHandler.UpdateProject)
	// // Delete one Project
	project.Delete("/:projectId", projectHandler.DeleteProject)
}
