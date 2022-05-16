package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server listen on port 8000")
	http.ListenAndServe(":8000", nil)
}
