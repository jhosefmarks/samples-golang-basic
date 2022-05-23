package main

import (
	"fmt"
	"net/http"

	"http-server/routes"
)

func main() {
	routes.LoadRoutes()

	fmt.Println("Server listen on port 8000")

	http.ListenAndServe(":8000", nil)
}
