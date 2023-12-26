package projectHandler

import (
	"github.com/brianeno/projects-fiber/database"
	"github.com/brianeno/projects-fiber/internals/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetProjectsfunc gets all existing projects
// @Description Get all existing projects
// @Tags Projects
// @Accept json
// @Produce json
// @Success 200 {array} model.Project
// @router /api/project [get]
func GetProjects(c *fiber.Ctx) error {
	db := database.DB
	var projects []model.Project

	// find all projects in the database
	db.Find(&projects)

	// If no project is present return an error
	if len(projects) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No projects present", "data": nil})
	}

	// Else return projects
	return c.JSON(fiber.Map{"status": "success", "message": "Projects Found", "data": projects})
}

// CreateProject func create a project
// @Description Create a Project
// @Tags Project
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param category body string true "Category"
// @Param description body string true "Description"
// @Success 200 {object} model.Project
// @router /api/project [post]
func CreateProject(c *fiber.Ctx) error {
	db := database.DB
	project := new(model.Project)

	// Store the body in the project and return error if encountered
	err := c.BodyParser(project)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the project
	project.ID = uuid.New()
	// Create the Project and return error if encountered
	err = db.Create(&project).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create project", "data": err})
	}

	// Return the created project
	return c.JSON(fiber.Map{"status": "success", "message": "Created Project", "data": project})
}

// GetProject func one projects by ID
// @Description Get one project by ID
// @Tags Project
// @Accept json
// @Produce json
// @Success 200 {object} model.Project
// @router /api/project/{id} [get]
func GetProject(c *fiber.Ctx) error {
	db := database.DB
	var project model.Project

	// Read the param projectId
	id := c.Params("projectId")

	// Find the project with the given Id
	db.Find(&project, "id = ?", id)

	// If no such project present return an error
	if project.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No project present", "data": nil})
	}

	// Return the project with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Projects Found", "data": project})
}

// UpdateProject update a project by ID
// @Description Update a Project by ID
// @Tags Project
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param category body string true "Category"
// @Param description body string true "Description"
// @Success 200 {object} model.Project
// @router /api/project/{id} [post]
func UpdateProject(c *fiber.Ctx) error {
	type updateProject struct {
		Title       string `json:"title"`
		Category    string `json:"category"`
		Description string `json:"description"`
	}
	db := database.DB
	var project model.Project

	// Read the param projectId
	id := c.Params("projectId")

	// Find the project with the given Id
	db.Find(&project, "id = ?", id)

	// If no such project present return an error
	if project.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No project present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateProjectData updateProject
	err := c.BodyParser(&updateProjectData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the project
	project.Title = updateProjectData.Title
	project.Category = updateProjectData.Category
	project.Description = updateProjectData.Description

	// Save the Changes
	db.Save(&project)

	// Return the updated project
	return c.JSON(fiber.Map{"status": "success", "message": "Projects Found", "data": project})
}

// DeleteProject delete a Project by ID
// @Description Delete a project by ID
// @Tags Project
// @Accept json
// @Produce json
// @Success 200
// @router /api/project/{id} [delete]
func DeleteProject(c *fiber.Ctx) error {
	db := database.DB
	var project model.Project

	// Read the param projectId
	id := c.Params("projectId")

	// Find the project with the given Id
	db.Find(&project, "id = ?", id)

	// If no such project present return an error
	if project.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No project present", "data": nil})
	}

	// Delete the project and return error if encountered
	err := db.Delete(&project, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete project", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted project"})
}
