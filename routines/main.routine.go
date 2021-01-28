package routines

import (
	"fmt"
	"golang_kafka/app/routes"
	"golang_kafka/config"
	"golang_kafka/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// InitMain function fiber
func InitMain() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	routes.ProductRoutes(app)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}
