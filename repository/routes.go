package repository

import "github.com/gofiber/fiber/v2"

func (repo *Repository) SetupRoutes(app *fiber.App) {
	//routes.Setup(repo, app)
	app.Static("/", "./client/public")

	api := app.Group("/api")
	api.Get("/alerts", repo.GetAlerts)
	api.Post("/alerts", repo.CreateAlert)
	api.Patch("/alerts/:id", repo.UpdateAlert)
	api.Delete("/alerts/:id", repo.DeleteAlert)
	api.Get("/alerts/:id", repo.GetAlertByID)
}
