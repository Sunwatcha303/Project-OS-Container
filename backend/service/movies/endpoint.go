package movies

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
	api_key := c.GetHeader("Api-Key")
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

func (e *Endpoint) GetAllMovie(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[movie] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	if response, err := e.logic.GetAllMovieLogic(); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[movie] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
	} else if len(*response) == 0 {
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func (e *Endpoint) GetMovieById(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[movie] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	id := c.Param("movie_id")
	if response, err := e.logic.GetMovieByIdLogic(id); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[movie] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func (e *Endpoint) GetScoreBymovieid(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[movie] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	id := c.Param("movie_id")
	if response, err := e.logic.GetScoreBymovieidLogic(id); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[movie] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func (e *Endpoint) AddMovie(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[movie] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	var requestBody MovieRequest
	if err := c.BindJSON(&requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.InternalServerError)
		fmt.Printf("[movie] Internal Server error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	if err := e.logic.AddMovieLogic(requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.InternalServerError)
		fmt.Printf("[movie] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (e *Endpoint) UpdateMoviebyId(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[movie] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	id := c.Param("movie_id")
	var requestBody MovieUpdateRequest
	if err := c.BindJSON(&requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.InternalServerError)
		fmt.Printf("[movie] Internal Server error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	if err := e.logic.UpdateMoviebyIdLogic(id, requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[movie] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (e *Endpoint) DeleteMoviebyId(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[movie] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	id := c.Param("movie_id")
	if err := e.logic.DeleteMoviebyIdLogic(id); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[movie] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, nil)
}
