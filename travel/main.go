package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/saga-pattern-study/travel/controllers/flightsearch"
	"github.com/jonathanmorais/saga-pattern-study/travel/controllers/healthchecker"
	"saga-pattern-study/travel/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToMongo()
}

func main() {
	router := gin.New()
	router.POST("/health'", healthchecker.Ok)
	router.POST("/flightsaga", flightsearch.FlightSearch)
	// router.GET("flightcheck", flightforward.FlightConsumer)
	router.Run(":8090")

}
