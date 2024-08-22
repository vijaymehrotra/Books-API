package routes

import (
	"log"
	"vijaymehrotra/go-deploy/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes() {
	app := fiber.New()
	app.Static("/", "./static")

    // Serve the static HTML file
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile("./static/home.html")
    })

	app.Get("/books", controller.GetBooks)
	app.Post("/create_book", controller.CreateBook)
	app.Delete("/delete_book/:id", controller.DeleteBook)
	app.Put("/update_book/:id", controller.UpdateBook)

	log.Fatal(app.Listen(":3000"))
}