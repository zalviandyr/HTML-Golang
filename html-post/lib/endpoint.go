package lib

import (
	"database/sql"
	"net/http"
	"path"
	"text/template"
)

type Endpoint struct {
	DB *sql.DB
}

func (e Endpoint) GetRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var basePath = path.Join("view", "base.html")
		var registerPath = path.Join("view", "register.html")

		var template = template.Must(template.New("").ParseFiles(registerPath, basePath))

		var data = map[string]string{
			"title": "Register",
		}

		CheckErr(template.ExecuteTemplate(w, "base", data))
	}
}

func (e Endpoint) PostRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var basePath = path.Join("view", "base.html")
		var submitPath = path.Join("view", "submit.html")

		var template = template.Must(template.New("").ParseFiles(submitPath, basePath))

		// parse to form
		CheckErr(r.ParseForm())

		// get data from form
		var email = r.FormValue("email")
		var username = r.FormValue("username")
		var firstName = r.FormValue("firstName")
		var lastName = r.FormValue("lastName")
		var password, salt = TextToSha1(r.FormValue("password"))

		//insert into database
		sql := "INSERT INTO users(email, first_name, last_name, username, password, salt) VALUES(?,?,?,?,?,?)"
		var stmt, err = e.DB.Prepare(sql)
		CheckErr(err)
		stmt.Exec(email, firstName, lastName, username, password, salt)

		var data = map[string]string{
			"title":    "Submit",
			"status":   "Success insert into database",
			"email":    email,
			"username": username,
			"fullName": firstName + " " + lastName,
			"password": password,
		}

		CheckErr(template.ExecuteTemplate(w, "base", data))
	}
}
