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

		p.Id = id
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

func DeleteProduct(id string) {
	db := db.ConnectDB()

	delete, err := db.Prepare("DELETE FROM PRODUCTS WHERE ID = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDB()

	update, err := db.Query(`
		SELECT ID, NAME, DESCRIPTION, PRICE, AMOUNT
		FROM PRODUCTS
		WHERE ID = $1
	`, id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for update.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = update.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount
	}

	defer db.Close()

	return product
}

func UpdateProduct(id int, name, description string, price float64, amount int) {
	db := db.ConnectDB()

	query, err := db.Prepare(`
		UPDATE PRODUCTS
			SET NAME = $1,
				DESCRIPTION = $2,
  			PRICE = $3,
				AMOUNT = $4
			WHERE ID = $5
	`)
	if err != nil {
		panic(err.Error())
	}

	query.Exec(name, description, price, amount, id)

	defer db.Close()
}
