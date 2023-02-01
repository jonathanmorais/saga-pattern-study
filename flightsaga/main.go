package main

import (
	"context"
	"log"
	"time"
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
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)

	}

	
	// conver
	message, err := json.Marshal(request)
    if err != nil {
        panic (err)
    }
	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: message},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
		// router := gin.New()
	// router.POST("/flightproducer", FlightProducer)
	// router.Run(":8070")
}

func main() {
	router := gin.New()
	router.POST("/flightproducer", FlightProducer)
	router.Run(":8070")

}
