package main

import (
	"database/sql"
	"log"
	"net/http"
	"path"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Employee struct {
	EmployeeID string
	LastName   string
	FirstName  string
	Title      string
	Address    string
	City       string
	Country    string
	HomePhone  string
}

func index(w http.ResponseWriter, r *http.Request) {

	sql := "SELECT EmployeeID, LastName, FirstName, Title, Address, City, Country, HomePhone FROM employees"

	var rows, err1 = db.Query(sql)
	checkErr(err1)
	defer rows.Close()

	var employees []map[string]string
	for rows.Next() {
		var employee Employee
		rows.Scan(
			&employee.EmployeeID, &employee.LastName, &employee.FirstName,
			&employee.Title, &employee.Address, &employee.City, &employee.Country,
			&employee.HomePhone,
		)

		var employeeMap = map[string]string{
			"employeeID": employee.EmployeeID,
			"fullName":   employee.FirstName + " " + employee.LastName,
			"address":    employee.Address + ", " + employee.City + ", " + employee.Country,
			"phone":      employee.HomePhone,
		}

		employees = append(employees, employeeMap)
	}

	var filePath = path.Join("view", "index.html")
	var template, err2 = template.ParseFiles(filePath)
	checkErr(err2)
	checkErr(template.Execute(w, employees))
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/northwind")
	checkErr(err)

	// http
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
