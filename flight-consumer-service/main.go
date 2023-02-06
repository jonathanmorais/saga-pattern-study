package main


import (
	"github.com/gin-gonic/gin"
	"saga-pattern-study/flight-consumer-service/consumer"
//	"saga-pattern-study/flight-consumer-service/producer"

)

func main() {
	router := gin.New()
	go consumer.SendForwardProducer()
	router.Run(":8060" )

}