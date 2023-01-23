package router

import (
	"github.com/gofiber/fiber"
	"github.com/saurabhkanawade/book_gorm_golang_project/controller"
	"log"
)

func Router() {
	log.Println("Starting the server on 9090..........")

	route := fiber.New()

	route.Get("/test", controller.Test)
	route.Get("/api/v1/books", controller.GetBooks)
	route.Get("/api/v1/books/:id", controller.GetBook)
	route.Post("/api/v1/books", controller.CreateBooks)
	route.Put("/api/v1/books/:id", controller.UpdateBook)
	route.Delete("/api/v1/books/:id", controller.DeleteBook)

	route.Listen(":9090")

}
