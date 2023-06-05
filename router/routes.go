package router

import (
	"net/http"
	controllers "weather/controller"
)

var WeatherRoutes = Routes{
	Route{"Get temp of city", http.MethodGet, "/getweather", controllers.GetCityWeather},
}
