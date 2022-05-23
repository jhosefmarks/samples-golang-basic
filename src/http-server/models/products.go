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

func InsertProduct(name, description string, price float64, amount int) {
	db := db.ConnectDB()

	insert, err := db.Prepare(`
		INSERT INTO PRODUCTS (
			NAME, DESCRIPTION, PRICE, AMOUNT
		) VALUES (
			$1, $2, $3, $4
		)
	`)
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, amount)
	defer db.Close()
}
