package routines

import (
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
			fmt.Println(data)
		}
	}
}
