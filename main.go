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
	Key string `json:"key"`
	Err string `json:"err"`
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

	err = validate.Name(fvr.Name)
	if err != nil {
		responseError(w, "name", err)
		return
	}

	err = validate.Name(fvr.SurName)
	if err != nil {
		responseError(w, "surName", err)
		return
	}

	err = validate.Email(fvr.Email)
	if err != nil {
		responseError(w, "email", err)
		return
	}

	err = validate.Password(fvr.Password)
	if err != nil {
		responseError(w, "password", err)
		return
	}

}


func responseError(w http.ResponseWriter, key string, err error)  {

	res := jsonResponseError{key, err.Error()}

	var b []byte
	b, err = json.Marshal(res)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error: %s", err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	w.Write(b)
}


func main() {

	fs := http.FileServer(http.Dir("ui"))
	http.Handle("/", fs)
	http.HandleFunc("/main_page", registration)

	log.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
