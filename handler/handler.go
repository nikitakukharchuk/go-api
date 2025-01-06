package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-api/database"
	"go-api/model"
)

func GetAllPosts(c *fiber.Ctx) error {
	db := database.DB.Db
	var posts []model.Post
	if err := db.Order("created_at DESC").Find(&posts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not retrieve posts",
			"data":    nil,
		})
	}
	if len(posts) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Posts not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Posts Found", "data": posts})
}

func GetPost(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var post model.Post
	db.Find(&post, "id = ?", id)
	if post.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Post not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Post Found", "data": post})
}

func CreatePost(c *fiber.Ctx) error {
	db := database.DB.Db
	post := new(model.Post)
	err := c.BodyParser(post)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&post).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create post", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Post has created", "data": post})
}

func DeletePost(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var post model.Post
	db.Find(&post, "id = ?", id)
	if post.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Post not found", "data": nil})
	}
	err := db.Delete(&post, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete post", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Post deleted"})

}

func UpdatePost(c *fiber.Ctx) error {
	type updatePost struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	db := database.DB.Db
	id := c.Params("id")
	var post model.Post
	db.Find(&post, "id = ?", id)

	if post.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Post not found", "data": nil})
	}

	var updateUserData updatePost
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	post.Title = updateUserData.Title
	post.Body = updateUserData.Body
	db.Save(&post)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Post updated !", "data": post})

}
