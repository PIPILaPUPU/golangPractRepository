package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Product struct {
	Product string
}

func (p Product) String() string {
	return fmt.Sprintf("Product: %s", p.Product)
}

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT product FROM products WHERE price > :price", sql.Named("price", 600))
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}

		err := rows.Scan(&product.Product)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(product)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return
	}
}
