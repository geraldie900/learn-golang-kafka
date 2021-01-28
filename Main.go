package main

import (
	"fmt"
	"golang_kafka/routines"
)

func main() {
	go routines.InitMain()
	go routines.InitKafkaReader()

	var input string
	fmt.Scanln(&input)
}
