package main

import (
	"database/sql"
	"html/html-post/lib"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// init database
	var db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/northwind")
	lib.CheckErr(err)

	// init endpoint
	var endpoint = lib.Endpoint{DB: db}

	// set url
	http.HandleFunc("/", endpoint.GetRegister)
	http.HandleFunc("/submit", endpoint.PostRegister)

	// start server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
