package models

import (
	"http-server/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func GetAllProducts() []Product {
	db := db.ConnectDB()

	rows, err := db.Query(`
		SELECT ID, NAME, DESCRIPTION, PRICE, AMOUNT
		FROM PRODUCTS
		ORDER BY NAME
	`)
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

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

	defer db.Close()
	return products
}
