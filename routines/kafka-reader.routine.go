package routines

import (
	"encoding/json"
	"fmt"
	"golang_kafka/app/functions"
	"golang_kafka/app/types"
)

// InitKafkaReader kafka reader router
func InitKafkaReader() {
	// consume kafka :
	kafkamsg := make(chan interface{})
	go functions.ConsumeData("products", kafkamsg)
	var data = types.KafkaMessage{}
	for {
		select {
		case resp := <-kafkamsg:
			data = resp.(types.KafkaMessage)
			var response types.Product
			err := json.Unmarshal(data.Value.Product, &response)
			// fmt.Println("BEFORE", data)
			if err != nil {
				//err
			}
			// fmt.Println(response)
			response2, err := json.Marshal(response)
			fmt.Println(string(response2))
		}
	}
}
