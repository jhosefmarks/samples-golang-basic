package controllers

import (
	"html/template"
	"net/http"

	"http-server/db"
	"http-server/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()

	rows, err := db.Query(`
		SELECT ID, NAME, DESCRIPTION, PRICE, AMOUNT
		FROM PRODUCTS
		ORDER BY NAME
	`)
	if err != nil {
		panic(err.Error())
	}

	p := models.Product{}
	products := []models.Product{}

	for rows.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
