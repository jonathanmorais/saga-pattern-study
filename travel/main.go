package main

import (
	"time"
	"os"
	"strconv"
	"fmt"
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/saga-pattern-study/travel/controllers/flightsearch"
	"github.com/jonathanmorais/saga-pattern-study/travel/controllers/healthchecker"
	"github.com/jonathanmorais/saga-pattern-study/travel/initializers"
	"github.com/jonathanmorais/saga-pattern-study/travel/pkg/memory_cache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToMongo()
}

func main() {
	// Logger
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	// use memory cache package
	c := memory_cache.GetInstance()
	// Readiness Probe Mock Config
	probeTimeRaw := os.Getenv("READINESS_PROBE_MOCK_TIME_IN_SECONDS")
	if probeTimeRaw == "" {
		probeTimeRaw = "5"
	}
	probeTime, err := strconv.ParseUint(probeTimeRaw, 10, 64)
	if err != nil {
		fmt.Println("Environment variable READINESS_PROBE_MOCK_TIME_IN_SECONDS conversion error", err)
	}
	c.Set("readiness.ok", "false", time.Duration(probeTime)*time.Second)

	// Creates a gin router with default middleware:
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

	router = gin.New()
	router.POST("/health'", healthchecker.Ok)
	router.POST("/flightsaga", flightsearch.FlightSearch)
	// router.GET("flightcheck", flightforward.FlightConsumer)
	router.Run(":8090")

}
