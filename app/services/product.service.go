package services

import (
	"encoding/json"
	"fmt"
	"golang_kafka/app/functions"
	"golang_kafka/app/types"
	"golang_kafka/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

// IntroduceProduct ....
func IntroduceProduct(c *fiber.Ctx) error {
	product := new(types.Product)

	if err := utils.ParseBodyAndValidate(c, product); err != nil {
		return err
	}

	product.ReleaseDate = time.Now().Format(time.RFC850)
	// fmt.Println("BEFORE : ", product)
	byteProduct, err := json.Marshal(product)
	// fmt.Println("AFTER :", byteProduct)
	if err != nil {
		fmt.Println("ERROR : json.marshal service", err)
	}

	kafkamessage := types.KafkaMessage{
		Value: types.KafkaValue{
			RequestType: "new product",
			Product:     byteProduct,
		},
	}

	sent := functions.ProduceData("products", kafkamessage)

	if sent {
		return c.Status(200).JSON(fiber.Map{
			"response_code": 200,
			"message":       "Product produced successfully",
		})
	}

	return nil
}
