package router

import (
	"log"
	"socialmedia/handlers"
	"socialmedia/middlewares"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

func Listen() {
	app := fiber.New()

	// Creates fiber-monitor instance
	prometheus := fiberprometheus.New("socialmedia")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Get("/logout", handlers.Logout, middlewares.Auth)

	post := app.Group("/post", middlewares.Auth)
	post.Post("/comment", handlers.Comment)
	post.Post("/like", handlers.Like)
	post.Get("/get/:id", handlers.Post)

	user := app.Group("/user", middlewares.Auth)
	user.Post("/update-pfp", handlers.ProfileUpdate)
	user.Post("/friend", handlers.Friend)

	userPost := user.Group("/post")
	userPost.Post("/create", handlers.CreatePost)
	userPost.Post("/edit", handlers.EditPost)
	userPost.Post("/delete", handlers.DeletePost)

	log.Fatal(app.Listen(":8080"))
}
