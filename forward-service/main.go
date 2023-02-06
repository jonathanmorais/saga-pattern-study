package main


import (
	"log"
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/gin-gonic/gin"

)

type Message struct{
	value string
}

// const (
// 	topic     = "flightfoward"
// 	partition = 0
// 	broker    = "localhost"
// 	port	  = ":9092"
// )


func FlightConsumer() string{
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "flightfoward",
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

func main() {
	router := gin.New()
	go FlightConsumer()
	router.Run(":8050")

}