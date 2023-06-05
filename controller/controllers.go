package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather/helper"

	"github.com/gin-gonic/gin"
)

func GetCityWeather(g *gin.Context) {

	city := g.Query("city")

	if city == "" {
		fmt.Println("City cant be empty")
		return

	}

	result, err := helper.CallWeatherApi(city)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occured",
		})
		return
	}

	defer result.Body.Close()
	var d map[string]interface{}

	err = json.NewDecoder(result.Body).Decode(&d)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Error occured",
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"city":         city,
		"weather_data": d,
	})

}
