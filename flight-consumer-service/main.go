package main


import (
	"log"
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/gin-gonic/gin"

)

const (
	topic     = "flighttopic"
	partition = 0
	broker    = "localhost"
	port	  = ":9092"
)

func FlightConsumer(c *gin.Context) {
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "flighttopic",
		GroupID: "g1",
		MaxBytes: 10,
	}
	
	reader := kafka.NewReader(conf)
	
	for {
		m, err := reader.ReadMessage(context.Background())
		log.Println("aqui1")
		if err != nil {
			log.Println("something wrong is happen: ", err)
			continue
		}
		log.Println("aqui1")
		log.Println("message consumed is: ", string(m.Value))
	}	

}

func main() {
	router := gin.New()
	router.POST("/flightconsumer", FlightConsumer)
	router.Run(":8060" )

}