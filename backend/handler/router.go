package handler

import (
	"net/http"

	"github.com/Sunwatcha303/Project-OS-Container/middleware"
	"github.com/Sunwatcha303/Project-OS-Container/service/central"
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
}

func (r Routes) InitRouter() http.Handler {
	centralEndpoint := central.NewEndpoint()

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
			Patten:      "/movie/",
			Endpoint:    centralEndpoint.Health,
			Validation:  middleware.NoMiddlewareValitdation,
		},
		// {
		// 	Name:        "POST: add movie",
		// 	Description: "ADD moive in the server",
		// 	Method:      http.MethodPost,
		// 	Patten:      "/add",
		// 	Endpoint:    nil,
		// 	Validation:  middleware.NoMiddlewareValitdation,
		// },
	}

	ro := gin.Default()
	ro.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Accept", "Context-type"},
	}))

	mainRoute := ro.Group("project-os-container")
	for _, e := range r.centralService {
		mainRoute.Handle(e.Method, e.Patten, e.Validation, e.Endpoint)
	}
	for _, e := range r.movieService {
		mainRoute.Handle(e.Method, e.Patten, e.Validation, e.Endpoint)
	}
	return ro
}
