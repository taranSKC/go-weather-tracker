package helper

import (
	"fmt"
	"net/http"
	"os"
)

func CallWeatherApi(city string) (*http.Response, error) {

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?" + "APPID=" + os.Getenv("API") + "&q=" + city)
	fmt.Println("http://api.openweathermap.org/data/2.5/weather?" + "APPID=" + os.Getenv("API") + "&q=" + city)

	return resp, err
}
