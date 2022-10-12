package main

// inserting records into a PostgreSQL database with Go's database/sql package
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

	// insert a row
	sqlStatement := `INSERT INTO cars (brand, model, city, year, price, description) 
    VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(sqlStatement, "Toyota", "4RUNNER", "Astana", "2007", "12000000", "вездеход")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
}
