package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Hello struct {
	Message string
}

type Name struct {
	FirstName string
	LastName  string
}

type Age struct {
	Age      int
	YearBorn int
}

type NameAge struct {
	Name Name
	Age  Age
}

var (
	netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	ageServiceUrl string
	logger        *log.Logger
)

func main() {
	ageServiceUrl = os.Getenv("AGE_SERVICE_URL")
	listenPort := os.Getenv("NOMAD_PORT_http")

	netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	logger = log.New(os.Stderr, "", 0)

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/name", nameHandler)
	mux.HandleFunc("/nameage", nameageHandler)
	http.ListenAndServe(":"+listenPort, mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data := Hello{"Hello World"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		logger.Fatal(err)
	}
	w.Write(resp)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	data := Name{"John", "Smith"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		logger.Fatal(err)
	}
	w.Write(resp)
}

func nameageHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/age", ageServiceUrl), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Fatal(fmt.Printf("Creating request error: %s", err.Error()))
		return
	}
	result, err := netClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Fatal(fmt.Printf("Do request error: %s", err.Error()))
		w.Write(fmt.Printf("Do request error: %s", err.Error()))
		return
	}

	var actual Age
	bodyBytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Fatal(fmt.Printf("Read body error: %s", err.Error()))
		return
	}
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
		w.WriteHeader(http.StatusInternalServerError)
		logger.Fatal(err)
		return
	}
	w.Write(resp)
}
