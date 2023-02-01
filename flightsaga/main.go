package main

import (
	"context"
	"log"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"


)

type Request struct {
	Message string `json:"message"`
}

func FlightProducer(c *gin.Context) {
	var request Request
	err := c.ShouldBindBodyWith(&request, binding.JSON); 
	if err != nil {
		log.Println("error: ", err)
	}

	// to produce messages
	topic := "flighttopic"
	//partition := 0

	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:   topic,
		Balancer: &kafka.LeastBytes{},
	}
		
	// conver
	out, err := json.Marshal(request)
    if err != nil {
        panic (err)
    }
	message := string(out) 

	err = w.WriteMessages(context.Background(),
	kafka.Message{
		Key:   []byte("Message"),
		Value: []byte(message),
	},)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	log.Println("Sended")

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func main() {
	router := gin.New()
	router.POST("/flightproducer", FlightProducer)
	router.Run(":8070")

}
