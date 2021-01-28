package functions

import (
	"encoding/json"
	"fmt"
	"golang_kafka/app/types"
	"golang_kafka/config/globals"

	"github.com/segmentio/kafka-go"
)

// ProduceData ....
func ProduceData(topic string, message types.KafkaMessage) bool {
	write := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	val, err := json.Marshal(message.Value)
	if err != nil {
		fmt.Println("ERROR : json.Marshal : ", err)
		return false
	}

	err = write.WriteMessages(globals.Ctx,
		kafka.Message{
			Key:   []byte(message.Key),
			Value: val,
		},
	)

	if err != nil {
		fmt.Println("ERROR : ", err)
		return false
	}

	if err = write.Close(); err != nil {
		fmt.Println("ERROR : ", err)
		return false
	}

	return true
}

// ConsumeData ....
func ConsumeData(topic string, message chan interface{}) {
	read := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
	})

	// var data = map[string]interface{}
	for {
		msg, err := read.ReadMessage(globals.Ctx)
		if err != nil {
			fmt.Println(err)
			break
		}
		kafkamsg := types.KafkaMessage{
			Key: string(msg.Key),
		}

		err = json.Unmarshal(msg.Value, &kafkamsg.Value)
		if err != nil {
			fmt.Println("ERROR : ", err)
			break
		}
		message <- kafkamsg

		//respjson, err := json.Marshal(kafkamsg)
		//message <- respjson
	}

	if err := read.Close(); err != nil {
		fmt.Println("ERROR", err)
		//return false
	}

	//return true
}
