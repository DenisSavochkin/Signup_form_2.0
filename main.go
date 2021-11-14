package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"server/validate"
)

func registration(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("cannot parse form. Error: %s", err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	passwordError := validate.Password(r.FormValue("password"))

	tmpl, _ := template.ParseFiles("ui/index.html")
	tmpl.Execute(w, passwordError)

}

func main() {

	fs := http.FileServer(http.Dir("ui"))
	http.Handle("/", fs)
	http.HandleFunc("/main_page", registration)

	log.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
