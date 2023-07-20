package main

import (
	"github.com/zeimedee/test2/internal/router"
)

func main() {

	port := ":8080"

	router := router.SetupRouter()

	router.Run(port)
}
