package router

import (
	"github.com/gofiber/fiber/v2"
	"go-api/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running")
	})

	api := app.Group("/api")
	postsApi := api.Group("/posts")
	postsApi.Get("/", handler.GetAllPosts)
	postsApi.Get("/:id", handler.GetPost)
	postsApi.Post("/", handler.CreatePost)
}
