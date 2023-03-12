package routes

import (
	"log"

	"mongo_gin/handlers"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	r := gin.Default()

	// Create blog post
	r.POST("/posts", handlers.CreateBlogPost)

	// Get blog post by ID
	r.GET("/posts/:id", handlers.GetBlogPost)

	// Get all blog posts
	r.GET("/posts", handlers.GetBlogPosts)

	// Update blog post by ID
	r.PUT("/posts/:id", handlers.UpdateBlogPost)

	// Delete blog post by ID
	r.DELETE("/posts/:id", handlers.DeleteBlogPost)

	log.Fatal(r.Run(":8080"))
}
