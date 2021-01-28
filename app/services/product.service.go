package services

import (
	"golang_kafka/app/functions"
	"golang_kafka/app/types"
	"golang_kafka/utils"

	"github.com/gofiber/fiber/v2"
)

// IntroduceProduct ....
func IntroduceProduct(c *fiber.Ctx) error {
	product := new(types.Product)

	if err := utils.ParseBodyAndValidate(c, product); err != nil {
		return err
	}

	kafkamessage := types.KafkaMessage{
		Value: types.KafkaValue{
			RequestType: "new product",
			Product:     *product,
		},
	}

	sent := functions.ProduceData("products", kafkamessage)

	if sent {
		return c.Status(200).JSON(fiber.Map{
			"message":       "Product produced successfully",
			"response_code": 200,
		})
	}

	return nil
}
