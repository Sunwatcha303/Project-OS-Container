package reviews

import (
	"fmt"

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

func (e *Endpoint) GetAllReviews(c *gin.Context) {

}

func (e *Endpoint) AddReview(c *gin.Context) {
	api_key := c.GetHeader("Api_Key")
	if api_key != constants.Api_Key {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.ApiKeyError)
		fmt.Printf("[movie] api key not found \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	var requestBody ReviewRequest
	if err := c.BindJSON(&requestBody); err != nil {
		httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.InternalServerError)
		fmt.Printf("[movie] Internal Server error \n%+v\n", errorResponse)
		c.AbortWithStatusJSON(httpStatusCode, errorResponse)
		return
	}
	// if err := /* */; err != nil {
	// 	httpStatusCode, errorResponse := templateError.GetErrorResponse(templateError.InternalServerError)
	// 	fmt.Printf("[movie] error \n%+v\n", errorResponse)
	// 	c.AbortWithStatusJSON(httpStatusCode, errorResponse)
	// 	return
	// }
}
