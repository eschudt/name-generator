package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/eschudt/name-generator/client"
)

type Hello struct {
	Message string
}

type Name struct {
	FirstName string
	LastName  string
}

type Age struct {
	Age int
}

type NameAge struct {
	Name Name
	Age  Age
}

var (
	netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	consulClient = client.NewConsulClient("http://127.0.0.1", "8500", "")
)

func init() {

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/name", nameHandler)
	mux.HandleFunc("/nameage", nameageHandler)
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

func nameageHandler(w http.ResponseWriter, r *http.Request) {
	address, port := consulClient.GetBaseURL("age-generator")
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s:%d/age", address, port), nil)
	result, _ := netClient.Do(req)

	var actual Age
	bodyBytes, _ := ioutil.ReadAll(result.Body)
	json.Unmarshal(bodyBytes, &actual)

	data := NameAge{
		Name: Name{
			FirstName: "John",
			LastName:  "Smith",
		},
		Age: actual,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}
