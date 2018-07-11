package main

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Message string
}

type Name struct {
	FirstName string
	LastName  string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/name", nameHandler)
	http.ListenAndServe(":8080", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data := Hello{"Hello World"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	data := Name{"John", "Smith"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}
