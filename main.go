package main

import (
	"fmt"
	"net/http"
	// "os"
	"io/ioutil"
	"database.sql"
	_ "github.com/lib/pq"
)

var INDEX_HTML []byte
var db *sql.DB

func main() {
	fmt.Println("Creating a web server here: http.localhost:3000")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/redirect", CreateRedirectHandler)
	http.ListenAndServe(":3000", nil)	
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(INDEX_HTML)
}

func CreateRedirectHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Creating Redirect Handler", r.Form)
}

func init() {
	INDEX_HTML, _ = ioutil.ReadFile("./templates/index.html")
	connectToDB()
}

func connectToDB() {
	var err error
	db, err := sql.Open("postgres", "host=/var/run/postgresql dbname=redirectapp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
}