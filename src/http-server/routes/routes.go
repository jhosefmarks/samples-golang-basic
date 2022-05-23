package routes

import (
	"net/http"

	"http-server/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
