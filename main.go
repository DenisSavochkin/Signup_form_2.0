package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DenisSavochkin/Signup_form_2.0/validate"
)

type validationError struct {
	FieldName string `json:"fieldName"`
	ErrorMessage string `json:"errorMessage"`
}

type jsonResponseError struct {
	Errors []validationError `json:"errors"`
}

type formValidationRequest struct {
	FirstName string `json:"name"`
	Surname   string `json:"surName"`
	Email     string `json:"email"`
	Password string `json:"password"`
}

func registration(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error: %s", err.Error()),
			http.StatusInternalServerError,
		)
	}
	defer r.Body.Close()

	fvr := &formValidationRequest{}
	err = json.Unmarshal(b, fvr)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error: %s", err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	jre := &jsonResponseError{
		Errors: make([]validationError, 0),
	}

	err = validate.Name(fvr.FirstName)
	if err != nil {
		jre.Errors = append(jre.Errors, validationError{
			FieldName: "FirstName",
			ErrorMessage: err.Error(),
		})
	}

	err = validate.Name(fvr.Surname)
	if err != nil {
		jre.Errors = append(jre.Errors, validationError{
			FieldName: "Surname",
			ErrorMessage: err.Error(),
		})
	}

	err = validate.Email(fvr.Email)
	if err != nil {
		jre.Errors = append(jre.Errors, validationError{
			FieldName: "Email",
			ErrorMessage: err.Error(),
		})
	}

	err = validate.Password(fvr.Password)
	if err != nil {
		jre.Errors = append(jre.Errors, validationError{
			FieldName: "Password",
			ErrorMessage: err.Error(),
		})
	}

	b, err = json.Marshal(jre)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error: %s", err.Error()),
			http.StatusInternalServerError,
		)
	}

	w.Write(b)
	fmt.Println(jre)

}


func main() {

	fs := http.FileServer(http.Dir("ui"))
	http.Handle("/", fs)
	http.HandleFunc("/validate", registration)

	log.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
