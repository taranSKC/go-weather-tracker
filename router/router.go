package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name       string
	Method     string
	Path       string
	handleFunc func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}

type Routes []Route

func (r routes) Weather(g *gin.RouterGroup) {

	weatherRoutes := g.Group("/" + os.Getenv("Weather"))

	for _, route := range WeatherRoutes {

		switch route.Method {
		case http.MethodGet:
			weatherRoutes.GET(route.Path, route.handleFunc)
		case http.MethodPost:
			weatherRoutes.POST(route.Path, route.handleFunc)

		}

	}

}

func Routing() {

	r := routes{
		router: gin.Default(),
	}

	grouping := r.router.Group(os.Getenv("Version"))

	r.Weather(grouping)

	r.router.Run(":" + os.Getenv("PORT"))

}
