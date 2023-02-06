package consumer


import (
	"log"
	"context"

	"github.com/segmentio/kafka-go"


)

type Message struct{
	value string
}

const (
	topic     = "flighttopic"
	partition = 0
	broker    = "localhost"
	port	  = ":9092"
)


func FlightConsumer() string{
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "flighttopic",
		GroupID: "g1",
		MaxBytes: 10,
	}
	
	reader := kafka.NewReader(conf)
	

	m, err := reader.ReadMessage(context.Background())
	if err != nil {
		log.Println("something wrong is happen: ", err)
	}
	log.Println("message consumed is: ", string(m.Value))

	return string(m.Value)
}

func SendForwardProducer() {
	msg := FlightConsumer()
	if msg == "" {
		log.Fatal("Message is empty")
	}

	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:   "flightfoward",
		Balancer: &kafka.LeastBytes{},
	}
		
	err := w.WriteMessages(context.Background(),
	kafka.Message{
		Key:   []byte("Message"),
		Value: []byte(msg),
	},)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	log.Println("Message sended to topic: flightfoward ")

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
