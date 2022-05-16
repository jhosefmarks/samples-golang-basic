package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Server listen on port 8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Blue T-shirt", Description: "Beautiful T-shirt", Price: 39.9, Amount: 5},
		{"Shoes", "Comfortable shoes for hiking.", 89.9, 3},
		{"Phone", "The best phone for you.", 259.99, 2},
		{"Jeans", "Casual clothes for everyday life.", 29.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", products)
}
