package types

// KafkaMessage ....
type KafkaMessage struct {
	Key   string     `json:"key"`
	Value KafkaValue `json:"value"`
}

// KafkaValue ....
type KafkaValue struct {
	RequestType string  `json:"request_type"`
	Product     Product `json:"product"`
}

// Product ...
type Product struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	ReleaseDate string `json:"release_date"`
}
