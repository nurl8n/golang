package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "kolesa"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `INSERT INTO cars (brand, model, city, year, price, description) 
    VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(sqlStatement, "BMW", "745i", "Almaty", "2008", "4000000", "классика")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nНовое авто добавлено успешно")
	}
}
