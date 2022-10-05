package main

import (
	"database/sql"
	"fmt"
	"log"
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
	http.HandleFunc("/data", dataRecord)
	http.ListenAndServe(":8080", nil)
}
func dataRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ALL CARS IN KOLESA-COPY:\n")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(404), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT * FROM cars")

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	snbs := make([]sandbox, 0)

	for rows.Next() {
		snb := sandbox{}
		err := rows.Scan(&snb.id, &snb.brand, &snb.model, &snb.city, &snb.year, &snb.price, &snb.description)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		snbs = append(snbs, snb)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, snb := range snbs {
		fmt.Fprintf(w, "%d %s %s %s %d %d %s\n", snb.id, snb.brand, snb.model, snb.city, snb.year, snb.price, snb.description)
		//fmt.Fprintf(w, snb.brand)
	}
}
