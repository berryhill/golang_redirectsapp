package main

import (
	"fmt"
	"net/http"
	"log"
// 	"os"
	"io/ioutil"
// 	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	_ "github.com/lib/pq"
)

var INDEX_HTML []byte
var db *sql.DB

func main() {
	fmt.Println("Creating a web server here: http.localhost:3000")
	setup()
}

func setup() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/redirect", CreateRedirectHandler)
	http.ListenAndServe(":3000", nil)	
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
//		if r.Host == HOST {
	log.Println("/GET")
	w.Write(INDEX_HTML)
//		return
//		}
//		RedirectHandler(w, r)	
}

//func RedirectHandler(w http.ResponseWriter, r http.Request) {
//	var destination string
//	db.QueryRow("SELECT destination FROM redirects WHERE source = $1", r.Host).Scan(&destination)
//	http.Redirect(w, r, destination, http.StatusMovedPermanently)
//}

func CreateRedirectHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db.Exec("INSERT INTO redirects(source, destination) VALUES($1, $2)", r.Form["source"][0], r.Form["destination"][0])
	fmt.Println("Creating Redirect Handler", r.Form)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func init() {
	INDEX_HTML, _ = ioutil.ReadFile("./templates/index.html")
	connectToDB()
}

func connectToDB() {
	var err error
	db, err = sql.Open("postgres", "host=/var/run/postgresql dbname=redirectsapp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

//	age := 21
//	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
}
