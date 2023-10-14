package movie

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	logic *MovieLogic
}

func NewEndpoint() *Endpoint {
	logic := MovieLogic{}
	return &Endpoint{
		logic: logic.InitMovieLogic(),
	}
}

func (e *Endpoint) Health(c *gin.Context) {
	response := HealthResponse{
		Code:        200,
		Status:      "up",
		Description: nil,
		Message:     nil,
	}
	c.JSON(http.StatusOK, response)
}

func (e *Endpoint) AddMovie(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}
