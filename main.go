package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// import pq as postgres driver
	_ "github.com/lib/pq"
)

const (
	user     = "testuser"
	port     = 5432
	host     = "localhost"
	password = "testpass"
	dbname   = "mike_db"
)

type Car struct {
	ID   int    "json:id"
	Make string "json:string"
}

func init() {
	fmt.Println("running init, nothing here now")
}

func checkErr(err error) {
	// if error is not nill, log fatal error
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("hello!")

	// let's test playing with postgres!
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", user, password, dbname, host, port)
	db, err := sql.Open("postgres", connStr)
	checkErr(err)

	// test hitting database test table
	rows, err := db.Query("SELECT * FROM test_table;")
	fmt.Println(rows)

	// create car struct from DB rows!
	var mycar Car // empty car object
	rows, err = db.Query("SELECT * FROM cars;")
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&mycar.ID, &mycar.Make)
		checkErr(err)
		fmt.Println("Car is ", mycar.Make)
	}

	// create mux router
	router := mux.NewRouter()
	router.HandleFunc("/test", testFunc).Methods("GET")

	// run the router
	log.Fatal(http.ListenAndServe(":8000", router))
}

// handler for "/test"
func testFunc(http.ResponseWriter, *http.Request) {
	fmt.Println("you hit me!")
}
