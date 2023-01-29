package flightsaga

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

const (
	topic     = "flighttopic"
	partition = 0
	broker    = "localhost"
)

func BrokerConn() bool {

	conn, err := kafka.DialLeader(context.Background(), "tcp", broker+":9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	return true
}

func SagaProducer(c *gin.Context) {

	conn := BrokerConn()
	if conn == false {
		log.Fatal("failed to take connection with broker:", conn)
	}

	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP(broker+":9092", broker+":9093", broker+":9094"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Flight"),
			Value: []byte("San Francisco"),
		},
	)
	log.Println(w.Logger)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}

// func main() {
// 	router := gin.New()
// 	router.POST("/flightservice", SagaProducer)
// 	router.Run(":8070")

// }