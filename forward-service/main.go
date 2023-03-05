package main


import (
	"os"
	"log"
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	logi "github.com/rs/zerolog/log"
	"github.com/Depado/ginprom"


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
	// Logger
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logi.Logger = logi.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	router := gin.New()

	// Prometheus Exporter Config
	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	//Middlewares
	router.Use(p.Instrument())
	router.Use(gin.Recovery())
	
	
	go FlightConsumer()
	router.Run(":8050")

}