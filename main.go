package main

import (
	"encoding/json"
	"fmt"
	"github.com/DenisSavochkin/Signup_form_2.0/validate"
	"io/ioutil"
	"log"
	"net/http"
)

type formValidationRequest struct {
	Name string `json:"name"`
	SurName string `json:"surName"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type jsonResponseError struct {
	Errors []ValidationError
}

type ValidationError struct {
	FieldName string `json:"FieldName"`
	Message string `json:"Message"`
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

	var res jsonResponseError

	err = validate.Name(fvr.Name)
	if err != nil {
		res = responseError("name", err.Error(), res)
	}

	err = validate.Name(fvr.SurName)
	if err != nil {
		res = responseError("surName", err.Error(), res)
	}

	err = validate.Email(fvr.Email)
	if err != nil {
		res = responseError("email", err.Error(), res)
	}

	err = validate.Password(fvr.Password)
	if err != nil {
		res = responseError("password", err.Error(), res)
	}

	b, err = json.Marshal(res)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error: %s", err.Error()),
			http.StatusInternalServerError,
		)
	}

	w.Write(b)
	fmt.Println(res)

}


func responseError(fieldName, message string, res jsonResponseError) jsonResponseError  {

	res.Errors = append(res.Errors, ValidationError{fieldName, message})

	return res
}


func main() {

	fs := http.FileServer(http.Dir("ui"))
	http.Handle("/", fs)
	http.HandleFunc("/validate", registration)

	log.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
