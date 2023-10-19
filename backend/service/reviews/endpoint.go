package reviews

import (
	"fmt"
	"net/http"

	"github.com/Sunwatcha303/Project-OS-Container/constants"
	templateError "github.com/Sunwatcha303/Project-OS-Container/error"
	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	logic *ReviewsLogic
}

func NewEndpoint() *Endpoint {
	logic := ReviewsLogic{}
	return &Endpoint{
		logic: logic.InitReviewLogic(),
	}
}

func (e *Endpoint) Health(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[review] api key not found %+v\n", errorResponse)
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

func (e *Endpoint) GetAllReviews(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[review] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStusCode, errorResponse)
		return
	}
	if response, err := e.logic.GetAllReviewsLogic(); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[review] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
	} else if len(*response) == 0 {
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func (e *Endpoint) GetAllReviewsbyMovieId(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[review] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStusCode, errorResponse)
		return
	}
	id := c.Param("movie_id")
	if response, err := e.logic.GetAllReviewsbyMovieIdLogic(id); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[review] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
	} else if len(*response) == 0 {
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func (e *Endpoint) AddReview(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[review] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	var requestBody ReviewRequest
	if err := c.BindJSON(&requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.InternalServerError)
		fmt.Printf("[review] Internal Server error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	if err := e.logic.AddReviewLogic(requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[review] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (e *Endpoint) UpdateReviewbyId(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[review] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	id := c.Param("id")
	var requestBody ReviewUpdateRequest
	if err := c.BindJSON(&requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.InternalServerError)
		fmt.Printf("[review] Internal Server error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	if err := e.logic.UpdateReviewbyIdLogic(id, requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[review] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (e *Endpoint) DeleteReviewbyId(c *gin.Context) {
	api_key := c.GetHeader("Api-Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[review] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	id := c.Param("id")
	if err := e.logic.DeleteMoviebyIdLogic(id); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(err)
		fmt.Printf("[review] error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	c.JSON(http.StatusOK, nil)
}
