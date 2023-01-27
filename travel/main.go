package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/saga-pattern-study/travel/controllers/flightsearch"
	"github.com/jonathanmorais/saga-pattern-study/travel/controllers/healthchecker"
)

func main() {
	router := gin.New()
	router.POST("/health'", healthchecker.Ok)
	router.POST("/flightsaga", flightsearch.SagaProducer)
	router.Run()

}
