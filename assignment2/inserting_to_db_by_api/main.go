package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type sandbox struct {
	id          int
	brand       string
	model       string
	city        string
	year        int
	price       int
	description string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "kolesa"
)

func init() {
	connStr := "postgres://postgres:root@localhost/kolesa?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("connected to db")
}

func main() {
	http.HandleFunc("/insertdata", dataRecord)
	http.ListenAndServe(":8080", nil)
}
func dataRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Автомобиль добавлен\n")
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
	_, err = db.Exec(sqlStatement, "FORD", "F450", "Taraz", "2012", "15000000", "классика")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nНовое авто добавлено успешно")
	}
}
