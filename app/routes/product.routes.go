package routes

import (
	"golang_kafka/app/services"

	"github.com/gofiber/fiber/v2"
)

// ProductRoutes ...
func ProductRoutes(app fiber.Router) {
	r := app.Group("/product")

	r.Post("/create", services.IntroduceProduct)

}
