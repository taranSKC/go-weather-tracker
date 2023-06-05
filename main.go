package main

import (
	"fmt"
	"weather/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Env couldnt load")
		return
	}

	router.Routing()

}
