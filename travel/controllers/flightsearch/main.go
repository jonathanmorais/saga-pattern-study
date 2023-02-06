package flightsearch

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/gin-gonic/gin"

)

const flightServicePort = 8070

func FlightSearch(c *gin.Context) {
	requestURL := fmt.Sprintf("http://localhost:%d", flightServicePort)
	res, err := http.Get(requestURL)
	if err != nil {
		log.Println("error making http request: %s\n", err)
		os.Exit(1)

	}

	log.Println(res)
}
