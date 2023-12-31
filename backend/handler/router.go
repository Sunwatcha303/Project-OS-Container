package handler

import (
	"net/http"

	"github.com/Sunwatcha303/Project-OS-Container/middleware"
	"github.com/Sunwatcha303/Project-OS-Container/service/central"
	"github.com/Sunwatcha303/Project-OS-Container/service/movies"
	"github.com/Sunwatcha303/Project-OS-Container/service/reviews"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Patten      string
	Endpoint    gin.HandlerFunc
	Validation  gin.HandlerFunc
}
type Routes struct {
	router         *gin.Engine
	centralService []route
	movieService   []route
	reviewService  []route
}

func (r Routes) InitRouter() http.Handler {
	centralEndpoint := central.NewEndpoint()
	movieEndpoint := movies.NewEndpoint()
	reviewEndpoint := reviews.NewEndpoint()

	r.centralService = []route{
		{
			Name:        "Get: health",
			Description: "Get health status from server",
			Method:      http.MethodGet,
			Patten:      "/",
			Endpoint:    centralEndpoint.Health,
			Validation:  middleware.NoMiddlewareValitdation,
		},
	}

	r.movieService = []route{
		{
			Name:        "Get: health",
			Description: "Get health status from server",
			Method:      http.MethodGet,
			Patten:      "/movies/",
			Endpoint:    movieEndpoint.Health,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "GET: get all movie",
			Description: "Get all movie from server",
			Method:      http.MethodGet,
			Patten:      "/movies/all",
			Endpoint:    movieEndpoint.GetAllMovie,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "GET: get movie by id",
			Description: "Get movie from server by id",
			Method:      http.MethodGet,
			Patten:      "/movies/:movie_id",
			Endpoint:    movieEndpoint.GetMovieById,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "GET: get score by movie id",
			Description: "get score from server by id",
			Method:      http.MethodGet,
			Patten:      "/movies/score/:movie_id",
			Endpoint:    movieEndpoint.GetScoreBymovieid,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "POST: add movie",
			Description: "Add movie into server",
			Method:      http.MethodPost,
			Patten:      "/movies/add",
			Endpoint:    movieEndpoint.AddMovie,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "PUT: update movie by id",
			Description: "Update movie in server by id",
			Method:      http.MethodPut,
			Patten:      "/movies/update/:movie_id",
			Endpoint:    movieEndpoint.UpdateMoviebyId,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "DELETE: delete movie by id",
			Description: "Delete movie from server by id",
			Method:      http.MethodDelete,
			Patten:      "/movies/delete/:movie_id",
			Endpoint:    movieEndpoint.DeleteMoviebyId,
			Validation:  middleware.NoMiddlewareValitdation,
		},
	}

	r.reviewService = []route{
		{
			Name:        "Get: health",
			Description: "Get health status from server",
			Method:      http.MethodGet,
			Patten:      "/reviews/",
			Endpoint:    reviewEndpoint.Health,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "GET: get all reviews",
			Description: "Get all reviews from server",
			Method:      http.MethodGet,
			Patten:      "/reviews/all",
			Endpoint:    reviewEndpoint.GetAllReviews,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "GET: get all reviews by movie id",
			Description: "Get all review from server by movie id",
			Method:      http.MethodGet,
			Patten:      "/reviews/:movie_id",
			Endpoint:    reviewEndpoint.GetAllReviewsbyMovieId,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "POST: add review",
			Description: "Add review into server",
			Method:      http.MethodPost,
			Patten:      "/reviews/add",
			Endpoint:    reviewEndpoint.AddReview,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "PUT: update review by id",
			Description: "Update review from server by id",
			Method:      http.MethodPut,
			Patten:      "/reviews/update/:id",
			Endpoint:    reviewEndpoint.UpdateReviewbyId,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		{
			Name:        "DELETE: delete review by id",
			Description: "Delete review from server by id",
			Method:      http.MethodDelete,
			Patten:      "/reviews/delete/:id",
			Endpoint:    reviewEndpoint.DeleteReviewbyId,
			Validation:  middleware.NoMiddlewareValitdation,
		},
	}

	ro := gin.Default()
	ro.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Accept", "Api-key"},
	}))

	mainRoute := ro.Group("project-os-container")
	for _, e := range r.centralService {
		mainRoute.Handle(e.Method, e.Patten, e.Validation, e.Endpoint)
	}
	for _, e := range r.movieService {
		mainRoute.Handle(e.Method, e.Patten, e.Validation, e.Endpoint)
	}
	for _, e := range r.reviewService {
		mainRoute.Handle(e.Method, e.Patten, e.Validation, e.Endpoint)
	}
	return ro
}
