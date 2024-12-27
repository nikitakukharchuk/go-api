package router

import (
	"github.com/gofiber/fiber/v2"
	"go-api/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	postsApi := api.Group("/posts")
	postsApi.Get("/", handler.GetAllPosts)
	postsApi.Get("/:id", handler.GetPost)
	postsApi.Post("/", handler.CreatePost)
}
