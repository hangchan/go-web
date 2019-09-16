package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	username string = "testing"
	password string = "foo-foo-foo-foo"
	dbHost   string = "go-db"
	database string = "demo"
)

func counterHandler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"
	count := getCount()
	addCount(count)

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello from:  "+title+"\n")
	fmt.Fprintf(w, "Visitor Count: %d \n", count)
}

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"
	// test the function
	getCount()

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello from:  "+title+"\n")
}

func dbConn() string {
	dbConn := strings.Join([]string{username, ":", password, "@(", dbHost, ":3306)/", database, "?interpolateParams=true"}, "")

	return dbConn
}

func addCount(count int) {
	db, err := sql.Open("mysql", dbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	count++
	countStr := strconv.Itoa(count)

	sql := strings.Join([]string{"UPDATE COUNTER SET count = ", countStr}, "")

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

}

func getCount() int {
	db, err := sql.Open("mysql", dbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var count int
	sql := "SELECT COUNT FROM COUNTER"

	err = db.QueryRow(sql).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counterHandler)
	http.ListenAndServe(":8080", nil)
}
