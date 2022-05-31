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

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		idConvToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error in ID conversion:", err)
		}

		priceConvToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error in price conversion:", err)
		}

		amountConvToInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error in amount conversion:", err)
		}

		models.UpdateProduct(idConvToInt, name, description, priceConvToFloat, amountConvToInt)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
