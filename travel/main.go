package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/saga-pattern-study/travel/controllers/flightsaga"
	"github.com/jonathanmorais/saga-pattern-study/travel/controllers/healthchecker"
)

func main() {
	router := gin.New()
	router.POST("/health'", healthchecker.Ok)
	router.POST("/flightsaga", flightsaga.SagaProducer)
}
