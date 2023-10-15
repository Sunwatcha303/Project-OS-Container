package movie

import (
	"fmt"
	"net/http"

	"github.com/Sunwatcha303/Project-OS-Container/constants"
	templateError "github.com/Sunwatcha303/Project-OS-Container/error"
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
	api_key := c.GetHeader("Api_Key")
	fmt.Println("HelloWorld")
	fmt.Printf("test %+v", api_key)
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[movie] api key not found %+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	response := HealthResponse{
		Code:        200,
		Status:      "up",
		Description: nil,
		Message:     nil,
	}
	c.JSON(http.StatusOK, response)
}

func (e *Endpoint) AddMovie(c *gin.Context) {
	api_key := c.GetHeader("Api_Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[movie] api key not found %+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	var requestBody MovieRequest
	if err := c.BindJSON(&requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.InternalServerError)
		fmt.Printf("[movie] Internal Server error %+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	if err := e.logic.AddMovieLogic(requestBody); err != nil {

	}
	c.JSON(http.StatusCreated, nil)
}
