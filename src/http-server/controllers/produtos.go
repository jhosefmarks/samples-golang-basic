package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"http-server/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error in price conversion:", err)
		}

		amountConv, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error in amount conversion:", err)
		}

		models.InsertProduct(name, description, priceConv, amountConv)
	}

	http.Redirect(w, r, "/", 301)
}
