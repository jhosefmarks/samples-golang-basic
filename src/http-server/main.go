package main

import (
	"fmt"
	"net/http"

	"http-server/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)

	fmt.Println("Server listen on port 8000")
	http.ListenAndServe(":8000", nil)

}
