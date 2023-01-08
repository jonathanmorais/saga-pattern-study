package flightstep

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	city string `json:"city"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SendFlight(c *gin.Context) Response {
	var request Request
	err := c.BindJSON(&request)
	if err != nil {
		c.Err()
	}

	var response Response
	response.Status = http.StatusOK
	response.Message = request.city

	return response
}
