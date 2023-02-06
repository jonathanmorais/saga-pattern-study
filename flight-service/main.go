package main


import (
	"github.com/gin-gonic/gin"
	"saga-pattern-study/flight-service/consumer"
	// "saga-pattern-study/flight-service/producer"

)

func main() {
	router := gin.New()
	go consumer.SendForwardProducer()
	router.Run(":8060")

}